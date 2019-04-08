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

func importPkg(path, dir string) *build.Package {

	p, err := build.Import(path, dir, build.ImportComment)
	if err != nil {
		panic(err)
	}

	if p.BinaryOnly {
		return nil
	}

	if p.IsCommand() {
		return nil
	}

	return p
}

func extractTODO(fname string) {

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	cmap := ast.NewCommentMap(fset, f, f.Comments)
	for n, cgs := range cmap {
		f := fset.File(n.Pos())
		for _, cg := range cgs {
			t := cg.Text()
			if strings.Contains(t, pattern) {
				fmt.Printf("%s:%v:\n%s\n", fname, f.Position(cg.Pos()).Line, t)
			}
		}
	}
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	p := importPkg(os.Args[1], dir)
	for _, f := range p.GoFiles {
		extractTODO(filepath.Join(p.Dir, f))
	}
}
