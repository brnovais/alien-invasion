package config

import (
	"flag"
)

type Config struct {
	// You should create N aliens, where N is specified as a command-line argument.
	Aliens int

	// The map is in a file.
	MapFile string
}

type ConfigReader interface {
	Read() Config
}

type CommandLineConfigReader struct {
}

func (r CommandLineConfigReader) Read() Config {
	aliens := flag.Int("n", 0, "# of aliens")
	file := flag.String("file", "data.txt", "")

	flag.Parse()

	return Config{
		*aliens,
		*file,
	}
}
