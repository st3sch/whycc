package bankfile

type Converter interface {
	Comma() rune
	IsTransaction([]string) bool
	Convert([]string) []string
}
