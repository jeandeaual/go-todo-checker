package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	// Replace with `todo "../.."` if the repository is not in the GOPATH
	todo "github.com/m-rec/d508b714f416b2c7b0b935be70e04d17085cba2b"
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
	_, err = w.Write(responseBody)
	if err != nil {
		errorf("Couldn't write the error response body: %s\n", err)
	}
}

// httpHandler is the main request handler of the HTTP API
// @title TODO Checker
// @version 0.0.1
// @basepath /
// @host localhost
// @description Query Go package comments
// @license.name None
type httpHandler struct {
	workdir string
}

// @description Return the list of comments in the specified Go package containing a specific pattern.
// @id parse
// @produce json
// @param package path string true "Name of the package to parse"
// @param pattern query string false "Pattern to look for in the comments" default(TODO)
// @success 200 {object} todo.Comments
// @failure 400 {object} main.ErrorResponse
// @failure 404 {object} main.ErrorResponse
// @failure 500 {object} main.ErrorResponse
// @router /{package} [get]
func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// Only GET is supported
		replyWithError(http.StatusMethodNotAllowed, "Only the GET method is supported", w)
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
	comments := todo.NewComments()

	infof("Checking for %s in the comments of package %s\n", pattern, path)

	err := comments.Parse(path, h.workdir, pattern)
	if err != nil {
		message := fmt.Sprintf("Couldn't parse comments from package %s: %s\n", path, err)
		errorf(message)
		replyWithError(http.StatusBadRequest, message, w)
		return
	}

	infof("Found %d comment(s) in %s\n", len(comments), path)

	responseBody, err := json.Marshal(comments)
	if err != nil {
		// Marshalling a map of structs containing only serializable data types
		// cannot fail, but just in case, handle the error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(responseBody)
	if err != nil {
		errorf("Couldn't write the response body: %s\n", err)
	}
}
