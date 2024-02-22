package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHelloWorld(t *testing.T) {
	// Create a new Gin router
	r := gin.Default()

	// Define route and handler
	r.GET("/", HelloWorld)

	// Create a test HTTP server
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Make a GET request to the test server
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Error making GET request: %v", err)
	}
	defer res.Body.Close()

	// Check if the response status code is OK
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	// Read the response body
	var body string
	err = extractJSONResponse(res, &body)
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}

	// Check if the response body is "Hello World!"
	expectedBody := "Hello World!"
	if body != expectedBody {
		t.Errorf("Expected body %q, got %q", expectedBody, body)
	}
}

// Helper function to extract JSON response body
func extractJSONResponse(res *http.Response, v interface{}) error {
	decoder := json.NewDecoder(res.Body)
	return decoder.Decode(v)
}
