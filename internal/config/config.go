package config

import (
	"flag"
)

// Config represents avaliable options to configure game execution.
type Config struct {
	// You should create N aliens, where N is specified as a command-line argument.
	Aliens int

	// The program should run until each alien has moved at least 10,000 times.
	Iterations int

	// The map is in a file.
	MapFile string
}

// Reader contract representing how the configuration is read.
type Reader interface {
	Read() Config
}

// CommandLineReader implements the configuration
// reader from command-line arguments.
type CommandLineConfigReader struct {
}

// Read configuration from command-line options.
func (r CommandLineConfigReader) Read() Config {
	aliens := flag.Int("n", 0, "# of aliens")
	iterations := flag.Int("i", 1, "# of iterations")
	file := flag.String("file", "testdata/data.txt", "")

	flag.Parse()

	return Config{
		*aliens,
		*iterations,
		*file,
	}
}
