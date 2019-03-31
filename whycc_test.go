package main

import (
	"path"
	"reflect"
	"testing"
)

func TestFindCSVFilesInDirectory(t *testing.T) {

	testPath := path.Join("testdata", "outdir")

	filesToFind := []string{
		path.Join(testPath, "a.csv"),
		path.Join(testPath, "b.csv"),
	}

	filesFound, _ := findCSVFilesInDirectory(testPath)

	if !reflect.DeepEqual(filesToFind, filesFound) {
		t.Errorf("Result is '%#v' should be '%#v' for '%#v'", filesFound, filesToFind)
	}
}
