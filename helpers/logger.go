package helpers

import "github.com/rs/zerolog"

// SetupLogger setup logger system
func SetupLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
