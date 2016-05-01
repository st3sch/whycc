package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"bufio"
	"io"
	"log"
)

func main() {
	f,_ := os.Open("./testdata/Umsatzanzeige_1234567890_20160410.csv")
	r := csv.NewReader(bufio.NewReader(f))
	r.Comma =';'
	r.FieldsPerRecord = -1
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if len(record) != 9 || record[0] == "Buchung" {
			continue
		}

		fmt.Println(record)
	}
}



