package hotelbeds

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateXSignature(t *testing.T) {
	// Store the original environment variables
	originalAPIKey := os.Getenv("API_KEY")
	originalSecret := os.Getenv("SECRET")

	// Set temporary environment variables for testing
	testAPIKey := "test-api-key"
	testSecret := "test-secret"

	os.Setenv("API_KEY", testAPIKey)
	os.Setenv("SECRET", testSecret)

	// Call the function to generate the X-Signature
	result := GenerateXSignature()

	// Expected X-Signature calculation
	expectedSignature := testAPIKey + testSecret + strconv.Itoa(int(time.Now().Unix()))
	hash := sha256.New()
	hash.Write([]byte(expectedSignature))
	expectedSignature = hex.EncodeToString(hash.Sum(nil))

	// Assert that the generated X-Signature matches the expected result
	assert.Equal(t, expectedSignature, result)

	// Restore the original environment variables
	os.Setenv("API_KEY", originalAPIKey)
	os.Setenv("SECRET", originalSecret)
}
