package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	assert.Nil(t, err)

	handler.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code, "handler returned a wrong status code")

	// Make sure the response body is an array
	// The comments in the "fmt" package might change depending the version
	// of Go, so don't check the exact payload
	var response []interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Nil(t, err)
}

func TestHTTPHandlerWithPattern(t *testing.T) {
	handler, rr := prepareServer(t)

	req, err := http.NewRequest("GET", "/net/http?pattern=FIXME", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code, "handler returned a wrong status code")

	// Make sure the response body is an array
	// The comments in the "net/http" package might change depending the version
	// of Go, so don't check the exact payload
	var response []interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Nil(t, err)
}

func TestHTTPHandlerPOST(t *testing.T) {
	handler, rr := prepareServer(t)

	req, err := http.NewRequest("POST", "/fmt", nil)
	assert.Nil(t, err)

	handler.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code, "handler returned a wrong status code")
}

func TestHTTPHandlerNoPackage(t *testing.T) {
	handler, rr := prepareServer(t)

	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	handler.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusBadRequest, rr.Code, "handler returned a wrong status code")
}

func TestHTTPHandlerInvalidPackage(t *testing.T) {
	handler, rr := prepareServer(t)

	req, err := http.NewRequest("GET", "/invalid/package", nil)
	assert.Nil(t, err)

	handler.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusBadRequest, rr.Code, "handler returned a wrong status code")
}
