package todo

import (
	"os"
	"testing"
)

func assertEqual(t *testing.T, expected interface{}, value interface{}, msg string) bool {
	if value != expected {
		t.Errorf("%s: got %v but expected %v", msg, value, expected)
		return false
	}

	return true
}

func TestImportPkg(t *testing.T) {
	// Standard package
	pkg, err := importPkg("fmt", "")
	if err != nil {
		t.Fatal("err should be nil")
	}
	assertEqual(t, "fmt", pkg.Name, "the wrong package was imported")

	// Command package
	_, err = importPkg("cmd/go", "")
	if err == nil {
		t.Error(err)
	}

	// Non-existing package
	_, err = importPkg("", "")
	if err == nil {
		t.Error(err)
	}
}

func TestExtractPattern(t *testing.T) {
	// TODO: Comment used for testing

	comments, err := extractPattern("parser_test.go", "TODO")
	if err != nil {
		// If no comment was extracted, don't execute the following tests
		t.Fatal(err)
	}
	if !assertEqual(t, 1, len(comments), "retrieved an invalid number of comments") {
		return
	}
	assertEqual(t, "parser_test.go", comments[0].Filename, "invalid filename")
	assertEqual(t, "TODO: Comment used for testing\n", comments[0].Text, "invalid comment text")
}

func TestCommentsParse(t *testing.T) {
	var comments Comments

	workdir, _ := os.Getwd()

	err := comments.Parse("fmt", workdir, "TODO")
	if err != nil {
		t.Fatal("err should be nil")
	}

	// The implementation is already tested by TestImportPkg and TestExtractPattern
}
