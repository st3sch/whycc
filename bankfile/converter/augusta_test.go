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
