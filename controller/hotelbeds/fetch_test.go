// hotelbeds_test.go
package hotelbeds

import (
	"github.com/stretchr/testify/assert"
	. "nuiteApi/models"
	"testing"
)

func TestFormatHotelBedsResponse_Success(t *testing.T) {
	// Create a sample JSON response for testing
	jsonResponse := []byte(`{
		"hotels": {
			"hotels": [
				{"Code": 123, "Currency": "USD", "MinRate": "150.5"},
				{"Code": 456, "Currency": "EUR", "MinRate": "200.75"}
			]
		}
	}`)

	// Call the function to be tested
	result, err := formatHotelBedsResponse(jsonResponse)

	// Assert that the function behaves as expected
	assert.NoError(t, err)

	// You can add more assertions based on your expected result
	expectedResult := HotelsResult{
		Data: []struct {
			HotelId  string  `json:"hotelId"`
			Currency string  `json:"currency"`
			Price    float64 `json:"price"`
		}{
			{"123", "USD", 150.5},
			{"456", "EUR", 200.75},
		},
	}

	assert.Equal(t, expectedResult, result)
}

func TestFormatHotelBedsResponse_ErrorDecoding(t *testing.T) {
	// Create an invalid JSON response for testing
	invalidJSON := []byte(`{"invalid": "json"`)

	// Call the function to be tested
	result, err := formatHotelBedsResponse(invalidJSON)

	// Assert that the function returns an error
	assert.Error(t, err)
	assert.Nil(t, result)
}

// Add more test cases as needed to cover different scenarios

//func TestGetJson(t *testing.T) {
//
//	// Create a JSON payload for testing
//	jsonPayload := []byte(`{
//  "stay": {
//    "checkIn": "2024-06-15",
//    "checkOut": "2024-06-16"
//  },
//  "occupancies": [
//    {
//      "rooms": 1,
//      "adults": 2,
//      "children": 0
//    }
//  ],
//  "hotels": {
//    "hotel": [
//      77,
//      168,
//      264,
//      265,
//      297,
//      311
//    ]
//  }
//}`)
//
//	// Call the function to be tested
//	statusCode, _ := GetJson(jsonPayload)
//
//	assert.Equal(t, http.StatusOK, statusCode)
//	assert.Equal(t, statusCode, 200)
//
//}
