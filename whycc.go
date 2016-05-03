package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"st3sch/whycc/converter"
)

func main() {

	files, err := filepath.Glob("./testdata/Umsatzanzeige_1234567890_*.csv")
	if err != nil {
		log.Fatal(err)
	}
	for _, filename := range files {
		fmt.Println(filename)
	}

	f, err := os.Open("./testdata/Umsatzanzeige_1234567890_20160410.csv")
	if err != nil {
		log.Fatal(err)
	}

	err = ConvertFile(f, os.Stdout, converter.NewIngDiBa())
	if err != nil {
		log.Fatal(err)
	}
}

func ConvertFile(in io.Reader, out io.Writer, c converter.Converter) error {
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
