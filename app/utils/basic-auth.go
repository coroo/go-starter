package utils

import (
	"encoding/base64"
	"os"
)

func CreateAuth() string {
	auth := os.Getenv("BASIC_AUTH_USERNAME") + ":" + os.Getenv("BASIC_AUTH_PASSWORD")
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
