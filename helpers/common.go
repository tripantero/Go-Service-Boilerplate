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