package config

import (
	"flag"
)

type Config struct {
	// You should create N aliens, where N is specified as a command-line argument.
	Aliens int

	// The program should run until each alien has moved at least 10,000 times.
	Iterations int

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
	iterations := flag.Int("i", 1, "# of iterations")
	file := flag.String("file", "data.txt", "")

	flag.Parse()

	return Config{
		*aliens,
		*iterations,
		*file,
	}
}
