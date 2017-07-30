package bankfile

import (
	"reflect"
	"testing"

	"github.com/st3sch/whycc/bankfile/converter"
)

func TestFactory_FindBy(t *testing.T) {
	f := factory{}

	c, err := f.FindBy("ingdiba")
	if err != nil {
		t.Errorf("Unexpected error: '%v'", err)
	}
	_, ok := c.(Converter)
	if !ok {
		t.Errorf("Returned converter doesn't implement interface 'converter'")
	}

	_, err = f.FindBy("non existent")
	if err == nil {
		t.Error("Expected invlaid type error not raised")
	}
}

func TestFactory_loadConverterFor(t *testing.T) {
	type testpair struct {
		in       string
		expected reflect.Type
	}

	tests := []testpair{
		{"ingdiba", reflect.TypeOf(converter.IngDiBa{})},
		{"augusta", reflect.TypeOf(converter.Augusta{})},
		{"krspaka", reflect.TypeOf(converter.KrSpaKa{})},
	}

	f := factory{}
	for _, pair := range tests {
		conv, err := f.loadConverterFor(pair.in)
		if err != nil {
			t.Errorf("Unexpected Error: '%v'", err)
			continue
		}

		converterType := reflect.TypeOf(conv)
		if pair.expected != converterType {
			t.Errorf("Expected converter type %q does not match returned converter type %q", pair.expected, converterType)
		}
	}

	_, err := f.loadConverterFor("non existent")
	if err == nil {
		t.Error("Expected invlaid type error not raised")
	}
}

func TestNewConverterLocator(t *testing.T) {
	cl := NewConverterLocator()
	if reflect.TypeOf(cl) != reflect.TypeOf(factory{}) {
		t.Errorf("ConverterLocator does not return factory. It returns: %q", reflect.TypeOf(cl))
	}
}
