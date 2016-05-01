package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"bufio"
	"io"
	"log"
	"time"
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

		record = convert(record)
		fmt.Println(record)
	}
}

func convert(record []string) []string  {
	result := make([]string, 6)

	// Date
	t, err := time.Parse("02.01.2006", record[0])
	if err != nil {
		log.Fatal(err)
	}
	result[0] = t.Format("02/01/2006")

	// Payee
	result[1] = record[2]

	// Memo
	result[3] = record[4]

	// Amount
	amount := record[5]
	if(amount[0:1] == "-") {
		result[5] = amount[1:]
	} else {
		result[4] = amount
	}

	return result
}



