package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const defaultPattern = "TODO"

func main() {
	// Parse the command-line arguments
	var pattern string

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s PACKAGE\n\nFlags:\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}

	flag.StringVar(&pattern, "pattern", defaultPattern, "Pattern to search for in the package comments")

	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprint(os.Stderr, "A package is required\n\n")
		flag.Usage()
		os.Exit(1)
	}

	// Retrieve the current working directory
	workdir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't retrieve the current working directory: %s\n", err)
		os.Exit(1)
	}

	path := flag.Args()[0]

	// Parse the package comments
	var comments Comments

	err = comments.Parse(path, workdir, pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't parse comments from package: %s\n", err)
		os.Exit(1)
	}

	for _, comment := range comments {
		fmt.Printf("%s:%v:\n%s\n", comment.Filename, comment.Line, comment.Text)
	}
}
