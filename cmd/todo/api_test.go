package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
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

func prepareServer(t *testing.T) (http.Handler, *httptest.ResponseRecorder) {
	// Retrieve the current working directory
	workdir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	return &httpHandler{workdir: workdir}, httptest.NewRecorder()
}

func TestHTTPHandler(t *testing.T) {
	handler, rr := prepareServer(t)

	req, err := http.NewRequest("GET", "/fmt", nil)
	if err != nil {
		t.Fatal("err should be nil")
	}

	handler.ServeHTTP(rr, req)

	// Check the status code
	assertEqual(t, http.StatusOK, rr.Code, "handler returned a wrong status code")

	// Make sure the response body is an array
	// The comments in the "fmt" package might change depending the version
	// of Go, so don't check the exact payload
	var response []interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("err should be nil")
	}
}

func TestHTTPHandlerWithPattern(t *testing.T) {
	handler, rr := prepareServer(t)

	req, err := http.NewRequest("GET", "/net/http?pattern=FIXME", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, req)

	// Check the status code
	assertEqual(t, http.StatusOK, rr.Code, "handler returned a wrong status code")

	// Make sure the response body is an array
	// The comments in the "net/http" package might change depending the version
	// of Go, so don't check the exact payload
	var response []interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("err should be nil")
	}
}

func TestHTTPHandlerPOST(t *testing.T) {
	handler, rr := prepareServer(t)

	req, err := http.NewRequest("POST", "/fmt", nil)
	if err != nil {
		t.Fatal("err should be nil")
	}

	handler.ServeHTTP(rr, req)

	// Check the status code
	assertEqual(t, http.StatusMethodNotAllowed, rr.Code, "handler returned a wrong status code")
}

func TestHTTPHandlerNoPackage(t *testing.T) {
	handler, rr := prepareServer(t)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal("err should be nil")
	}

	handler.ServeHTTP(rr, req)

	// Check the status code
	assertEqual(t, http.StatusBadRequest, rr.Code, "handler returned a wrong status code")
}

func TestHTTPHandlerInvalidPackage(t *testing.T) {
	handler, rr := prepareServer(t)

	req, err := http.NewRequest("GET", "/invalid/package", nil)
	if err != nil {
		t.Fatal("err should be nil")
	}

	handler.ServeHTTP(rr, req)

	// Check the status code
	assertEqual(t, http.StatusBadRequest, rr.Code, "handler returned a wrong status code")
}
