package converter

import (
	"log"
)

type Augusta struct {
	comma rune
}

func NewAugusta() Augusta {
	return Augusta{
		comma: ';',
	}
}

func (a Augusta) Comma() rune {
	return a.comma
}

func (a Augusta) IsTransaction(record []string) bool {
	return !(len(record) != 13 || record[0] == "Buchungstag" || record[9] == "Anfangssaldo" || record[9] == "Endsaldo")
}

func (a Augusta) Convert(record []string) []string {
	result := make([]string, 6)
	var err error

	// Date
	result[0], err = convertDateFrom("02.01.2006", record[0])
	if err != nil {
		log.Fatal(err)
	}

	// Payee
	result[1] = record[3]

	// Memo
	result[3] = record[8]

	// Amount
	amount := convertThousandAndCommaSeparator(record[11])
	switch record[12] {
	case "H":
		result[5] = amount
	case "S":
		result[4] = amount
	default:
		log.Println("No SOLL or HABEN flag given!")
	}

	return result
}
