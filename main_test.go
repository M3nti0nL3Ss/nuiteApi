// main_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupRouter(t *testing.T) {
	router := setupRouter()

	// Create a mock HTTP request
	req, err := http.NewRequest("GET", "/hotels", nil)
	assert.NoError(t, err)

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Serve the request to the router
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusInternalServerError, w.Code)

}
