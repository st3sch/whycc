package converter

import (
	"reflect"
	"testing"
)

func TestKrSpaKa_Comma(t *testing.T) {
	expectedComma := ';'
	k := NewKrSpaKa()
	if k.Comma() != expectedComma {
		t.Errorf("Field seperator is %q instead of expected %q", k.Comma(), expectedComma)
	}
}

func TestKrSpaKa_IsTransaction(t *testing.T) {
	type testpair struct {
		expected bool
		record   []string
	}

	tests := []testpair{
		{false, []string{}},
		{false, []string{"1", "2", "3", "4", "5", "6", "7"}},
		{false, []string{"Auftragskonto", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17"}},
		{true, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17"}},
	}

	k := NewKrSpaKa()
	for _, pair := range tests {
		actual := k.IsTransaction(pair.record)
		if actual != pair.expected {
			t.Errorf("Result is '%v', but it should be '%v' for %q", actual, pair.expected, pair.record)
		}
	}
}

func TestKrSpaKa_Convert(t *testing.T) {
	type testpair struct {
		record   []string
		expected []string
	}

	tests := []testpair{
		{
			[]string{"123456789", "14.10.16", "14.10.16", "FOLGELASTSCHRIFT", "1212-3434 acommpany ", "XC352647657", "sfasfae-325678fsdgr", "2323462477", "", "", "", "ACOMPANY", "DE1234567891234", "YYYYAAAAA", "-12,99", "EUR", "Umsatz gebucht"},
			[]string{"14/10/2016", "ACOMPANY", "", "[FOLGELASTSCHRIFT] 1212-3434 acommpany ", "", "12.99"},
		}, {
			[]string{"123456789", "15.10.16", "16.10.16", "FOLGELASTSCHRIFT", "1212-3434 bcommpany ", "XC352647657", "sfasfae-325678fsdgr", "2323462477", "", "", "", "BCOMPANY", "DE1234567891234", "YYYYAAAAA", "3092,44", "EUR", "Umsatz gebucht"},
			[]string{"15/10/2016", "BCOMPANY", "", "[FOLGELASTSCHRIFT] 1212-3434 bcommpany ", "3092.44", ""},
		},
	}

	k := NewKrSpaKa()
	for _, pair := range tests {
		actual := k.Convert(pair.record)
		if !reflect.DeepEqual(actual, pair.expected) {
			t.Errorf("Actual record (%q) does't match expected record (%q) with input (%q)", actual, pair.expected, pair.record)
		}
	}
}
