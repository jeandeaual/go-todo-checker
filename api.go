package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// ErrorResponse is a struct representing the response sent by the API
// when an error occurs
type ErrorResponse struct {
	Message string `json:"message"`
}

// replyWithError writes an error to the response payload and appropriately sets
// the Content-Type and response code
func replyWithError(statusCode int, message string, w http.ResponseWriter) {
	responseBody, err := json.Marshal(ErrorResponse{
		Message: message,
	})
	if err != nil {
		// Marshalling a struct containing only serializable data types
		// cannot fail, but just in case, handle the error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseBody)
}

// httpHandler is the main request handler of the HTTP API
func httpHandler(workdir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			// Only GET is supported
			replyWithError(http.StatusBadRequest, "Only the GET method is supported", w)
			return
		}

		// Assume that anything after the first / in the URL is the package to parse
		path := strings.TrimPrefix(r.URL.Path, "/")

		if len(path) == 0 {
			replyWithError(http.StatusBadRequest, "You need to specify a package", w)
			return
		}

		queryValues := r.URL.Query()

		pattern := defaultPattern

		// Check if the user provided a pattern in the query string
		if param := queryValues.Get("pattern"); len(param) > 0 {
			pattern = param
		}

		// Parse the package comments
		comments := NewComments()

		log.Printf("Checking for %s in the comments of package %s\n", pattern, path)

		err := comments.Parse(path, workdir, pattern)
		if err != nil {
			message := fmt.Sprintf("Couldn't parse comments from package %s: %s\n", path, err)
			log.Println(message)
			replyWithError(http.StatusBadRequest, message, w)
			return
		}

		log.Printf("Found %d comment(s) in %s\n", len(comments), path)

		responseBody, err := json.Marshal(comments)
		if err != nil {
			// Marshalling a map of structs containing only serializable data types
			// cannot fail, but just in case, handle the error
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBody)
	}
}
