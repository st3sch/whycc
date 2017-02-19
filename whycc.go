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
	inputdir := flag.String("i", ".", "Input directory")
	outputdir := flag.String("o", ".", "Output directory")
	flag.Parse()
	fmt.Println("Inputdir: ", *inputdir)
	fmt.Println("Outputdir: ", *outputdir)

	converterLocator := bankfile.NewConverterLocator()
	for banktype, pattern := range patterns {
		fmt.Println("Banktype: ", banktype)
		fmt.Println("Pattern: ", *pattern)
		if *pattern == "" {
			continue
		}

		files, err := filepath.Glob(*inputdir + string(filepath.Separator) + *pattern)
		fmt.Println(files)
		if err != nil {
			log.Fatal(err)
		}

		conv, err := converterLocator.FindBy(banktype)
		if err != nil {
			log.Fatal(err)
		}

		for _, filename := range files {
			fmt.Println("File: ", filename)
			f, err := os.Open(filename)
			if err != nil {
				log.Fatal(err)
			}

			outfilename := *outputdir + string(filepath.Separator) + banktype + "_" + path.Base(filename)
			outfile, err := os.Create(outfilename)
			if err != nil {
				log.Fatal(err)
			}
			defer outfile.Close()

			err = ConvertFile(f, outfile, conv)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func ConvertFile(in io.Reader, out io.Writer, c bankfile.Converter) error {
	r := csv.NewReader(in)
	r.Comma = c.Comma()
	r.FieldsPerRecord = -1

	w := csv.NewWriter(out)

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
