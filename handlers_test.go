package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testTransaction struct {
	Payer     string `json:"payer"`
	Points    int    `json:"points"`
	Timestamp string `json:"timestamp"`
}

// Test that a GET request to balances returns the payer balances with
// the HTTP code 200
func TestGetBalances(t *testing.T) {
	router := getRouter()

	router.GET("/balances", getBalances)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/balances", nil)

	testHTTPResponse(t, router, req, func(response *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := response.Code == http.StatusOK
		p, err := ioutil.ReadAll(response.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}

// Test that a POST request to transactions returns the new transaction with
// the HTTP code 200
func TestAddTransaction(t *testing.T) {
	router := getRouter()

	// Define the route similar to its definition in the routes file
	router.POST("/transactions", newTransaction)
	// Create a request to send to the above route
	body := testTransaction{Payer: "DANNON", Points: 1000, Timestamp: "2020-11-02T14:00:00Z"}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(body)
	req, _ := http.NewRequest("POST", "/transactions", payloadBuf)

	testHTTPResponse(t, router, req, func(response *httptest.ResponseRecorder) bool {
		// Test that the http status code is 201
		statusOK := response.Code == 201
		p, err := ioutil.ReadAll(response.Body)
		pageOK := err == nil && strings.Index(string(p), "added") > 0

		return statusOK && pageOK
	})
}

// Test that a POST request to spend returns the payer points used with
// the HTTP code 200
func TestSpendPoints(t *testing.T) {
	router := getRouter()

	// Define the route similar to its definition in the routes file
	router.POST("/transactions", newTransaction)

	body := testTransaction{Payer: "DANNON", Points: 1000, Timestamp: "2020-11-02T14:00:00Z"}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(body)
	// Create a request to send to the above route
	req, _ := http.NewRequest("POST", "/transactions", payloadBuf)

	testHTTPResponse(t, router, req, func(response *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := response.Code == http.StatusOK
		p, err := ioutil.ReadAll(response.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Article 1</title>") > 0

		return statusOK && pageOK
	})
}
