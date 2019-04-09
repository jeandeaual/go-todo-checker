package main

import (
	"fmt"
	"os"
)

func main() {
	workdir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't retrieve the current working directory: %s\n", err)
		os.Exit(1)
	}

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s PACKAGE\n", os.Args[0])
		os.Exit(1)
	}

	path := os.Args[1]

	var comments Comments

	err = comments.Parse(path, workdir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't parse comments from package: %s\n", err)
		os.Exit(1)
	}

	for _, comment := range comments {
		fmt.Printf("%s:%v:\n%s\n", comment.Filename, comment.Line, comment.Text)
	}
}
