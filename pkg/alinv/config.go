package alinv

import (
	"github.com/brnovais/alien-invasion/internal/config"
)

// ConfigReaderEnum represents available configuration readers.
// We need a single byte to represent all available options.
// Since this is a simulation, some values will not be implemented.
type ConfigReaderEnum byte

const (
	// CommandLineConfigReader is used to read
	// configuration from command-line arguments.
	CommandLineConfigReader ConfigReaderEnum = iota

	// EnvironmentConfigReader will be used to read
	// configuration from environment variables.
	EnvironmentConfigReader

	// FileConfigReader will be used to read configuration from files.
	FileConfigReader
)

// NewConfigReader creates a new configuration reader based on the type supplied
// as an argument. Since this is a simulation, only CommandLineConfigReader
// will be implemented.
func NewConfigReader(cfgType ConfigReaderEnum) config.Reader {
	switch cfgType {
	case CommandLineConfigReader:
		return config.CommandLineConfigReader{}
	default:
		return nil
	}
}
