package hotelbeds

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"strconv"
	"time"
)

func GenerateXSignature() string {
	xSignature := os.Getenv("API_KEY")
	xSignature += os.Getenv("SECRET")
	xSignature += strconv.Itoa(int(time.Now().Unix()))
	hash := sha256.New()
	hash.Write([]byte(xSignature))
	xSignature = hex.EncodeToString(hash.Sum(nil))

	return xSignature
}
