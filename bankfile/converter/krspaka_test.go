package converter

import (
	"testing"
)

func TestKrSpaKa_Comma(t *testing.T) {
	expectedComma := ';'
	k := NewKrSpaKa()
	if k.Comma() != expectedComma {
		t.Errorf("Field seperator is %q instead of expected %q", k.Comma(), expectedComma)
	}
}

type testpair struct {
	record   []string
	expected bool
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
