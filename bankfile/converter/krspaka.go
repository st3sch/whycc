package converter

import (
	"fmt"
	"log"
)

type KrSpaKa struct {
	comma rune
}

func NewKrSpaKa() KrSpaKa {
	return KrSpaKa{
		comma: ';',
	}
}

func (k KrSpaKa) Comma() rune {
	return k.comma
}

func (k KrSpaKa) IsTransaction(record []string) bool {
	return !(len(record) != 17 || record[0] == "Auftragskonto")
}

func (k KrSpaKa) Convert(record []string) []string {
	result := make([]string, 6)
	var err error

	// Date
	result[0], err = convertDateFrom("02.01.06", record[1])
	if err != nil {
		log.Fatal(err)
	}

	// Payee
	result[1] = record[11]

	// Memo
	result[3] = fmt.Sprintf("[%v] %v", record[3], record[4])

	return result
}
