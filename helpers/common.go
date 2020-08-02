package helpers

import "os"

// GetPORT function to validate the port selection
func GetPORT() string {
	port := os.Getenv("PORT")
	if len(port) <= 4 {
		port = ":6007"
	}
	return port
}

// PanicIfErrNotNil i don't know this function is worth to exists or not
func PanicIfErrNotNil(err error) {
	if err != nil {
		panic(err.Error())
		os.Exit(0)
	}
}
