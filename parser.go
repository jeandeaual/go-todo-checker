package main

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// Comments represents a set of Go file comments
type Comments []*Comment

// Comment represents a Go comment line
type Comment struct {
	Filename string
	Line     int
	Text     string
}

// Parse parses the comments contained in a Go package
func (c *Comments) Parse(path, dir, pattern string) error {
	pkg, err := importPkg(path, dir)
	if err != nil {
		return err
	}

	for _, file := range pkg.GoFiles {
		fname := filepath.Join(pkg.Dir, file)
		fileComments, err := extractPattern(fname, pattern)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't retrieve comments from %s: %s\n", fname, err)
			continue
		}

		*c = append(*c, fileComments...)
	}

	return nil
}

// importPkg imports a Go package from a path
func importPkg(path, dir string) (*build.Package, error) {
	pkg, err := build.Import(path, dir, build.ImportComment)
	if err != nil {
		return nil, err
	}

	// Don't parse binary-only packages (they don't contain any comment)
	if pkg.BinaryOnly {
		return nil, fmt.Errorf("package %s is binary-only", path)
	}

	// Don't parse command packages
	if pkg.IsCommand() {
		return nil, fmt.Errorf("package %s is a command", path)
	}

	return pkg, nil
}

// extractPattern extracts comments containing TODO from a Go source file
func extractPattern(fname, pattern string) ([]*Comment, error) {
	comments := []*Comment{}

	// Parse the file and create the AST
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)
	if err != nil {
		return comments, err
	}

	cmap := ast.NewCommentMap(fset, f, f.Comments)
	for n, cgs := range cmap {
		// Retrieve the file pointer
		file := fset.File(n.Pos())
		// Iterate through the comment groups
		for _, cg := range cgs {
			text := cg.Text()
			// Check if the comment group contains the pattern
			if strings.Contains(text, pattern) {
				comments = append(comments, &Comment{
					Filename: fname,
					Line:     file.Position(cg.Pos()).Line,
					Text:     text,
				})
			}
		}
	}

	return comments, nil
}
