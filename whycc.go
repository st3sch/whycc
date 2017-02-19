package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"path"

	"github.com/st3sch/whycc/bankfile"
)

func main() {
	patterns := make(map[string]*string)
	patterns["ingdiba"] = flag.String("ingdiba", "", "Pattern for ING DiDba files")
	patterns["augusta"] = flag.String("augusta", "", "Pattern for Augusta Bank files")
	patterns["krspaka"] = flag.String("krspaka", "", "Pattern for Kreissparkasse Augsburg files")
	inDir := flag.String("i", ".", "Input directory")
	outDir := flag.String("o", ".", "Output directory")
	cleanupInDir := flag.Bool("ci", false, "Delete input files after conversion")
	flag.Parse()
	fmt.Println("Inputdir: ", *inDir)
	fmt.Println("Outputdir: ", *outDir)

	converterLocator := bankfile.NewConverterLocator()
	for banktype, pattern := range patterns {
		fmt.Println("Banktype: ", banktype)
		fmt.Println("Pattern: ", *pattern)
		if *pattern == "" {
			continue
		}

		inFileNames, err := filepath.Glob(*inDir + string(filepath.Separator) + *pattern)
		fmt.Println(inFileNames)
		if err != nil {
			log.Fatal(err)
		}

		conv, err := converterLocator.FindBy(banktype)
		if err != nil {
			log.Fatal(err)
		}

		for _, inFileName := range inFileNames {
			fmt.Println("File: ", inFileName)
			inputFile, err := os.Open(inFileName)
			if err != nil {
				log.Fatal(err)
			}

			outFileName := *outDir + string(filepath.Separator) + banktype + "_" + path.Base(inFileName)
			outFile, err := os.Create(outFileName)
			if err != nil {
				log.Fatal(err)
			}
			defer outFile.Close()

			err = ConvertFile(inputFile, outFile, conv)
			if err != nil {
				log.Fatal(err)
			}

			if *cleanupInDir {
				fmt.Println("Delete " + inFileName)
				err := os.Remove(inFileName)
				if err != nil {
					log.Println("Could not delete file: " + inFileName)
					log.Println(err)
				}
			}
		}
	}
}

func ConvertFile(in io.Reader, out io.Writer, c bankfile.Converter) error {
	r := csv.NewReader(in)
	r.Comma = c.Comma()
	r.FieldsPerRecord = -1

	w := csv.NewWriter(out)
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
