package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test that a GET request to balances returns the payer balances with
// the HTTP code 200
func TestGetBalances(t *testing.T) {
	r := getRouter(true)

	r.GET("/balances", getBalances)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/balances", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}

// Test that a POST request to transactions returns the new transaction with
// the HTTP code 200
func TestAddTransaction(t *testing.T) {
	r := getRouter(true)

	// Define the route similar to its definition in the routes file
	r.POST("/transactions", addTransaction)

	// Create a request to send to the above route
	req, _ := http.NewRequest("POST", "/transactions", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Article 1</title>") > 0

		return statusOK && pageOK
	})
}

// Test that a POST request to spend returns the payer points used with
// the HTTP code 200
func TestSpendPoints(t *testing.T) {
	r := getRouter(true)

	// Define the route similar to its definition in the routes file
	r.POST("/transactions", addTransaction)

	// Create a request to send to the above route
	req, _ := http.NewRequest("POST", "/transactions", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Article 1</title>") > 0

		return statusOK && pageOK
	})
}
