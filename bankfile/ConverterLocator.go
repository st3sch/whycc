package bankfile

import "fmt"

type ConverterLocator interface {
	FindBy(string) (Converter, error)
}

type factory struct {
}

func (f factory) FindBy(bank string) (Converter, error) {
	return nil, fmt.Errorf("Invalid bank %v type", bank)
}

func NewConverterLocator() ConverterLocator {
	return &factory{}
}
