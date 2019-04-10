package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const defaultPattern = "TODO"

func main() {
	// Parse the command-line arguments
	var (
		serverMode bool
		pattern    string
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [PACKAGE]\n\nFlags:\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}

	flag.BoolVar(&serverMode, "server", false, "Server mode")
	flag.StringVar(&pattern, "pattern", defaultPattern, "Pattern to search for in the package comments")

	flag.Parse()

	// Retrieve the current working directory
	workdir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't retrieve the current working directory: %s\n", err)
		os.Exit(1)
	}

	if serverMode {
		// Expose an HTTP API
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			httpHandler(w, r, workdir)
		})
		log.Fatal(http.ListenAndServe(":8080", nil))
	}

	if flag.NArg() != 1 {
		fmt.Fprint(os.Stderr, "A package is required\n\n")
		flag.Usage()
		os.Exit(1)
	}

	// Standard command-line mode
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
