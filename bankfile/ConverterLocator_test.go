package bankfile

import (
	"reflect"
	"testing"
)

func TestFactory_FindBy(t *testing.T) {
	type testpair struct {
		in       string
		expected string
	}

	tests := []testpair{
		{"ingdiba", "IngDiBa"},
	}

	f := factory{}
	for _, pair := range tests {
		cl, err := f.FindBy(pair.in)
		if err != nil {
			t.Errorf("Unexpected Error: '%v'", err)
			continue
		}
		if pair.expected != reflect.TypeOf(cl).Name() {

		}
	}

	_, err := f.FindBy("non existent")
	if err == nil {
		t.Error("Expected invlaid type error not raised")
	}
}
