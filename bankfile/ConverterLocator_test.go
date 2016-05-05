package bankfile

import "testing"

func TestFactory_FindBy(t *testing.T) {

	f := factory{}
	_, err := f.FindBy("not existent")
	if err == nil {
		t.Error("Expected error not thrown")
	}
}
