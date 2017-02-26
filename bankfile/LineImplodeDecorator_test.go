package bankfile

import "testing"

type ConverterMock struct {
}

func (cm ConverterMock) Comma() rune {
	return ';'
}

func (cm ConverterMock) IsTransaction(record []string) bool {
	return false
}

func (cm ConverterMock) Convert(record []string) []string {
	return record
}

func TestLineImplodeDecorator_Comma(t *testing.T) {
	c := lineImplodeDecorator{ConverterMock{}}
	comma := c.Comma()
	if comma != ';' {
		t.Errorf("Comma value '%v' doesn't match expected ';'", comma)
	}
}

func TestLineImplodeDecorator_IsTransaction(t *testing.T) {
	c := lineImplodeDecorator{ConverterMock{}}
	record := []string{"aaa", "bbb"}
	result := c.IsTransaction(record)
	if result {
		t.Error("Expected false, got true")
	}
}

func TestLineImplodeDecorator_Convert(t *testing.T) {
	c := lineImplodeDecorator{ConverterMock{}}
	record := []string{"test\nline"}
	result := c.Convert(record)
	if result[0] != "test line" {
		t.Errorf("LineImplodeDacorator did not pass functin call to function 'Convert'")
	}
}
