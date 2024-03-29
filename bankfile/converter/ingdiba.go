package converter

import (
	"log"
)

type IngDiBa struct {
	validRecordLenght int
	indexOfDate       int
	indexOfPayee      int
	indexOfMemo       int
	indexOfAmount     int
	comma             rune
}

func NewIngDiBa() IngDiBa {
	return IngDiBa{
		validRecordLenght: 9,
		indexOfDate:       0,
		indexOfPayee:      2,
		indexOfMemo:       4,
		indexOfAmount:     7,
		comma:             ';',
	}
}

func (i IngDiBa) Comma() rune {
	return i.comma
}

func (i IngDiBa) IsTransaction(record []string) bool {
	return !(len(record) != i.validRecordLenght || record[0] == "Buchung")
}

func (i IngDiBa) Convert(record []string) []string {
	result := make([]string, 6)
	var err error

	// Date
	result[0], err = convertDateFrom("02.01.2006", record[i.indexOfDate])
	if err != nil {
		log.Fatal(err)
	}

	// Payee
	result[1] = record[i.indexOfPayee]

	// Memo
	result[3] = record[i.indexOfMemo]

	// Amount
	amount := convertThousandAndCommaSeparator(record[i.indexOfAmount])
	if isNegative(amount) {
		result[4] = abs(amount)
	} else {
		result[5] = amount
	}

	return result
}
