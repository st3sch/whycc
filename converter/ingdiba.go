package converter

import (
	"log"
	"time"
)

type IngDiBa struct {
	comma rune
}

func NewIngDiBa() *IngDiBa {
	return &IngDiBa{
		comma: ';',
	}
}

func (i *IngDiBa) Comma() rune {
	return i.comma
}

func (i *IngDiBa) IsTransaction(record []string) bool {
	return !(len(record) != 9 || record[0] == "Buchung")
}

func (i *IngDiBa) Convert(record []string) []string {
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
	if amount[0:1] == "-" {
		result[5] = amount[1:]
	} else {
		result[4] = amount
	}

	return result
}
