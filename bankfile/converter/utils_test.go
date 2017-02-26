package converter

import "testing"

func TestImplodeLines(t *testing.T) {
	type testpair struct {
		field    string
		expected string
	}

	tests := []testpair{
		{"hello\nworld", "hello world"},
		{"hello\r\nworld", "hello world"},
		{"hello\rworld", "helloworld"},
		{"\nhello\nworld\n", "hello world"},
	}

	for _, pair := range tests {
		actual := ImplodeLines(pair.field)
		if actual != pair.expected {
			t.Errorf("Result is '%#v' should be '%#v' for '%#v'", actual, pair.expected, pair.field)
		}
	}
}
