package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/st3sch/whycc/bankfile"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	patterns := make(map[string]*string)
	patterns["ingdiba"] = flag.String("ingdiba", "", "Pattern for ING DiDb files")
	inputdir := flag.String("i", ".", "Input directory")
	flag.Parse()

	converterLocator := bankfile.NewConverterLocator()
	for banktype, pattern := range patterns {
		fmt.Println("Banktype: ", banktype)
		fmt.Println("Pattern: ", *pattern)
		files, err := filepath.Glob(*inputdir + string(filepath.Separator) + *pattern)
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

			err = ConvertFile(f, os.Stdout, conv)
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
