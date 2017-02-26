package bankfile

import "github.com/st3sch/whycc/bankfile/converter"

type lineImplodeDecorator struct {
	bankfileConverter Converter
}

func (c lineImplodeDecorator) Comma() rune {
	return c.bankfileConverter.Comma()
}

func (c lineImplodeDecorator) IsTransaction(record []string) bool {
	return c.bankfileConverter.IsTransaction(record)
}

func (c lineImplodeDecorator) Convert(record []string) []string {
	record = c.bankfileConverter.Convert(record)
	for index, field := range record {
		record[index] = converter.ImplodeLines(field)
	}
	return record
}
