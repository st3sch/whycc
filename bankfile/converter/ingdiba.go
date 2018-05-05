package converter

import (
	"log"
)

type IngDiBa struct {
	comma rune
}

func NewIngDiBa() IngDiBa {
	return IngDiBa{
		comma: ';',
	}
}

func (i IngDiBa) Comma() rune {
	return i.comma
}

func (i IngDiBa) IsTransaction(record []string) bool {
	return !(len(record) != 9 || record[0] == "Buchung")
}

func (i IngDiBa) Convert(record []string) []string {
	result := make([]string, 6)
	var err error

	// Date
	result[0], err = convertDateFrom("02.01.2006", record[0])
	if err != nil {
		log.Fatal(err)
	}

	// Payee
	result[1] = record[2]

	// Memo
	result[3] = record[4]

	// Amount
	amount := convertThousandAndCommaSeparator(record[7])
	if isNegative(amount) {
		result[4] = abs(amount)
	} else {
		result[5] = amount
	}

	return result
}
