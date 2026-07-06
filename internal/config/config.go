package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Timestamp int
	Timezone  int
	Format    string
}

const DefaultFormat = "02.01.2006 15:04:05"

func Parse() (*Config, error) {
	cfg := &Config{
		Format: DefaultFormat,
	}

	flag.IntVar(&cfg.Timestamp, "t", 0, "Unix timestamp (required)")
	flag.IntVar(&cfg.Timezone, "tz", 0, "Timezone offset in hours (default: 0 = UTC)")
	flag.StringVar(&cfg.Format, "f", cfg.Format, "Output time format (Go layout)")

	flag.Usage = PrintUsage

	flag.Parse()

	if cfg.Timestamp == 0 {
		flag.Usage()
		return nil, fmt.Errorf("timestamp is required. Use -t <timestamp>")
	}

	return cfg, nil
}

func PrintUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "A simple flexible Unix timestamp converter")
	fmt.Fprintln(os.Stderr, "\nOptions:")
	flag.PrintDefaults()
}
