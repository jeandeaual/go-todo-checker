package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func prepareServer(t *testing.T) (http.HandlerFunc, *httptest.ResponseRecorder) {
	// Retrieve the current working directory
	workdir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	return http.HandlerFunc(httpHandler(workdir)), httptest.NewRecorder()
}

func TestHTTPHandler(t *testing.T) {
	handler, rr := prepareServer(t)

	req, err := http.NewRequest("GET", "/fmt", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned the wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Make sure the response body is an array
	// The comments in the "fmt" package might change depending the version
	// of Go, so don't check the exact payload
	var response []interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
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
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned the wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Make sure the response body is an array
	// The comments in the "net/http" package might change depending the version
	// of Go, so don't check the exact payload
	var response []interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}
}

func TestHTTPHandlerPOST(t *testing.T) {
	handler, rr := prepareServer(t)

	req, err := http.NewRequest("POST", "/fmt", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned the wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestHTTPHandlerInvalidPackage(t *testing.T) {
	handler, rr := prepareServer(t)

	req, err := http.NewRequest("GET", "/invalid/package", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned the wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
