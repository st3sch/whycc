package bankfile

import (
	"fmt"

	"github.com/st3sch/whycc/bankfile/converter"
)

type ConverterLocator interface {
	FindBy(string) (Converter, error)
}

type factory struct {
}

func (f factory) FindBy(bank string) (Converter, error) {
	switch bank {
	case "ingdiba":
		return converter.NewIngDiBa(), nil
	case "augusta":
		return converter.NewAugusta(), nil
	case "krspaka":
		return converter.NewKrSpaKa(), nil
	default:
		return nil, fmt.Errorf("Invalid bank %v type", bank)
	}
}

func NewConverterLocator() ConverterLocator {
	return factory{}
}
