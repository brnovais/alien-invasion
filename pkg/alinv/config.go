package alinv

import (
	"github.com/brnovais/alien-invasion/internal/config"
)

// We need a single byte to represent all available options.
type ConfigReaderEnum byte

const (
	// Read configuration from command-line arguments.
	CommandLineConfigReader ConfigReaderEnum = iota

	// Read configuration from environment variables.
	EnvironmentConfigReader

	// Read configuration from files.
	FileConfigReader
)

// Create a new configuration reader based on the type supplied as an argument.
// For now, only CommandLineConfigReader is implemented.
func NewConfigReader(cfgType ConfigReaderEnum) config.ConfigReader {
	switch cfgType {
	case CommandLineConfigReader:
		return config.CommandLineConfigReader{}
	default:
		return nil
	}
}
