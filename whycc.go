package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/st3sch/whycc/bankfile"
)

var (
	version string
	commit  string
	date    string
)

func main() {
	// Flag definition
	patterns := make(map[string]*string)
	patterns["ingdiba"] = flag.String("ingdiba", "", "Pattern for ING DiDba files")
	patterns["augusta"] = flag.String("augusta", "", "Pattern for Augusta Bank files")
	patterns["krspaka"] = flag.String("krspaka", "", "Pattern for Kreissparkasse Augsburg files")

	inDir := flag.String("i", ".", "Input directory")
	outDir := flag.String("o", ".", "Output directory")

	cleanupInDir := flag.Bool("ci", false, "Delete input files after conversion")
	cleanupOutDir := flag.Bool("co", false, "Delete all old csv files in output directory")

	printHelp := flag.Bool("h", false, "Print help")

	printVersion := flag.Bool("v", false, "Print version")

	flag.Parse()

	// print version
	if *printVersion {
		fmt.Println("Version:", version)
		fmt.Println("Commit: ", commit)
		fmt.Println("Date:   ", date)
	}

	// print help
	if *printHelp {
		flag.Usage()
	}

	// sanity checks
	if _, err := os.Stat(*inDir); os.IsNotExist(err) {
		log.Fatal("Input directory does not exist")
	}

	if _, err := os.Stat(*outDir); os.IsNotExist(err) {
		log.Fatal("Output directory does not exist")
	}

	// output dir cleanup
	if *cleanupOutDir {
		deleteAllCsvFilesInDirectory(*outDir)
	}

	// file parsing
	converterLocator := bankfile.NewConverterLocator()
	for banktype, pattern := range patterns {
		if *pattern == "" {
			continue
		}

		fmt.Println()
		fmt.Println("Parsing files of bank type: ", banktype)
		fmt.Println("-----------------------------------------------------------")

		inFileNames, err := filepath.Glob(*inDir + string(filepath.Separator) + *pattern)
		if err != nil {
			log.Fatal(err)
		}

		conv, err := converterLocator.FindBy(banktype)
		if err != nil {
			log.Fatal(err)
		}

		for _, inFileName := range inFileNames {
			fmt.Println("Converting file: ", inFileName)
			inputFile, err := os.Open(inFileName)
			if err != nil {
				log.Fatal(err)
			}

			outFileName := *outDir + string(filepath.Separator) + banktype + "_" + filepath.Base(inFileName)
			outFile, err := os.Create(outFileName)
			if err != nil {
				log.Fatal(err)
			}
			defer outFile.Close()

			err = ConvertFile(inputFile, outFile, conv)
			if err != nil {
				log.Fatal(err)
			}
			inputFile.Close()

			if *cleanupInDir {
				deleteFile(inFileName)
			}
		}
	}

	fmt.Println()
}

func ConvertFile(in io.Reader, out io.Writer, c bankfile.Converter) error {
	r := csv.NewReader(in)
	r.Comma = c.Comma()
	r.FieldsPerRecord = -1

	w := csv.NewWriter(out)
	w.UseCRLF = true
	w.Write([]string{"Date", "Payee", "Category", "Memo", "Outflow", "Inflow"})

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if !c.IsTransaction(record) {
			continue
		}

		record = c.Convert(record)
		err = w.Write(record)
		if err != nil {
			return err
		}
		w.Flush()
	}

	return nil
}

func deleteFile(fileName string) {
	fmt.Println("Deleting file: " + fileName)
	err := os.Remove(fileName)
	if err != nil {
		log.Println("Could not delete file: " + fileName)
		log.Println(err)
	}
}

func deleteAllCsvFilesInDirectory(dirName string) {
	fmt.Println("Clearing output directory ...")
	files, err := findCSVFilesInDirectory(dirName)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		deleteFile(file)
	}
}

func findCSVFilesInDirectory(dirName string) ([]string, error) {
	return filepath.Glob(dirName + string(filepath.Separator) + "*.[cC][sS][vV]")
}
