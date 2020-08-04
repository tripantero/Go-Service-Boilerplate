package helpers

import (
	"os"
	"strings"

	"github.com/rs/zerolog/log"
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

// Warn warned when error not too harmfull
func Warn(err error) {
	if err != nil {
		log.
			Warn().
			Msg(err.Error())
	}
}
