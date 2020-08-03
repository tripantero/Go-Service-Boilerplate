package helpers

import (
	"os"
	"strings"
)

// GetPORT function to validate the port selection
func GetPORT() string {
	port := os.Getenv("PORT")
	if !strings.Contains(port, ":") {
		port = ":" + port
	}
	if len(port) < 3 {
		port = ":6007"
	}
	return port
}

// PanicIfErrNotNil i don't know this function is worth to exists or not
func PanicIfErrNotNil(err error) {
	if err != nil {
		panic(err.Error())
	}
}
