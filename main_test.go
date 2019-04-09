package main

import (
	"testing"
)

func TestImportPkg(t *testing.T) {
	// Standard package
	pkg, err := importPkg("fmt", "")
	if err != nil {
		t.Error("err should be nil")
	}
	if pkg.Name != "fmt" {
		t.Error("the package name shouldn be \"fmt\"")
	}

	// Command package
	pkg, err = importPkg("cmd/go", "")
	if err == nil {
		t.Error("err shouldn't be nil")
	}

	// Non-existing package
	pkg, err = importPkg("", "")
	if err == nil {
		t.Error("err shouldn't be nil")
	}
}

func TestExtractTODO(t *testing.T) {
	// TODO: Comment used for testing

	err := extractTODO("main_test.go")
	if err != nil {
		t.Error("err should be nil")
	}
}
