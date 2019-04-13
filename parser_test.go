package todo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImportPkg(t *testing.T) {
	// Standard package
	pkg, err := importPkg("fmt", "")
	assert.Nil(t, err)
	assert.Equal(t, "fmt", pkg.Name)

	// Command package
	_, err = importPkg("cmd/go", "")
	assert.NotNil(t, err)

	// Non-existing package
	_, err = importPkg("", "")
	assert.NotNil(t, err)
}

func TestExtractPattern(t *testing.T) {
	// TODO: Comment used for testing

	comments, err := extractPattern("parser_test.go", "TODO")
	if !assert.Nil(t, err) {
		// If no comment was extracted, don't execute the following tests
		return
	}
	if !assert.Equal(t, 1, len(comments)) {
		return
	}
	assert.Equal(t, "parser_test.go", comments[0].Filename)
	assert.Equal(t, "TODO: Comment used for testing\n", comments[0].Text)
}

func TestCommentsParse(t *testing.T) {
	var comments Comments

	workdir, _ := os.Getwd()

	err := comments.Parse("fmt", workdir, "TODO")
	assert.Nil(t, err)

	// The implementation is already tested by TestImportPkg and TestExtractPattern
}
