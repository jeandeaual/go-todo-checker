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

const pattern = "TODO"

func importPkg(path, dir string) (*build.Package, error) {
	pkg, err := build.Import(path, dir, build.ImportComment)
	if err != nil {
		return nil, err
	}

	if pkg.BinaryOnly {
		return nil, fmt.Errorf("package %s is binary-only", path)
	}

	if pkg.IsCommand() {
		return nil, fmt.Errorf("package %s is a command", path)
	}

	return pkg, nil
}

func extractTODO(fname string) error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	cmap := ast.NewCommentMap(fset, f, f.Comments)
	for n, cgs := range cmap {
		file := fset.File(n.Pos())
		for _, cg := range cgs {
			text := cg.Text()
			if strings.Contains(text, pattern) {
				fmt.Printf("%s:%v:\n%s\n", fname, file.Position(cg.Pos()).Line, text)
			}
		}
	}

	return nil
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't retrieve the current working directory: %s\n", err)
		os.Exit(1)
	}

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s PACKAGE\n", os.Args[0])
		os.Exit(1)
	}

	pkgName := os.Args[1]

	pkg, err := importPkg(pkgName, dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't retrieve package: %s\n", err)
		os.Exit(1)
	}

	for _, file := range pkg.GoFiles {
		fname := filepath.Join(pkg.Dir, file)
		err := extractTODO(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't retrieve comments from %s: %s\n", fname, err)
		}
	}
}
