package bankfile

import (
	"reflect"
	"st3sch/whycc/bankfile/converter"
	"testing"
)

func TestFactory_FindBy(t *testing.T) {
	type testpair struct {
		in       string
		expected reflect.Type
	}

	tests := []testpair{
		{"ingdiba", reflect.TypeOf(converter.IngDiBa{})},
	}

	f := factory{}
	for _, pair := range tests {
		cl, err := f.FindBy(pair.in)
		if err != nil {
			t.Errorf("Unexpected Error: '%v'", err)
			continue
		}

		converterType := reflect.TypeOf(cl)
		if pair.expected != converterType {
			t.Errorf("Expected converter type %q does not match returned converter type %q", pair.expected, converterType)
		}
	}

	_, err := f.FindBy("non existent")
	if err == nil {
		t.Error("Expected invlaid type error not raised")
	}
}
