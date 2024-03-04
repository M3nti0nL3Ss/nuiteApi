// controller_test.go
package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetHotels(t *testing.T) {
	// Create a Gin router to handle the request
	router := gin.New()
	router.GET("/hotels", GetHotels)

	// Create a sample request URL with query parameters
	requestURL := "/hotels/?checkin=2024-06-15&checkout=2024-06-16&currency=USD&guestNationality=US&hotelIds=77,168,264,265,297,311&occupancies=[{\"rooms\": 1,\"adults\": 2,\"children\": 0}]"

	// Create a mock HTTP request using the GET method
	req, err := http.NewRequest("GET", requestURL, nil)
	assert.NoError(t, err)

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Serve the request to the router
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusMovedPermanently, w.Code)

}
