package converter

import (
	"testing"
)

func TestAugusta_Comma(t *testing.T) {
	expectedComma := ';'
	k := NewAugusta()
	if k.Comma() != expectedComma {
		t.Errorf("Field seperator is %q instead of expected %q", k.Comma(), expectedComma)
	}
}

func TestAugusta_IsTransaction(t *testing.T) {
	type testpair struct {
		expected bool
		record   []string
	}

	tests := []testpair{
		{false, []string{}},
		{false, []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}},
		{true, []string{"01.01.2020", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17"}},
	}

	k := NewAugusta()
	for _, pair := range tests {
		actual := k.IsTransaction(pair.record)
		if actual != pair.expected {
			t.Errorf("Result is '%v', but it should be '%v' for %q", actual, pair.expected, pair.record)
		}
	}
}
