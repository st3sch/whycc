package converter

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
	return result
}
