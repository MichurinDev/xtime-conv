package config

import (
	"flag"
	"os"
	"testing"
)

func resetFlags(args []string) {
	flag.CommandLine = flag.NewFlagSet(args[0], flag.ContinueOnError)
	os.Args = args
}

func TestParse_MissingTimestamp(t *testing.T) {
	resetFlags([]string{"xtime-conv"})

	_, err := Parse()
	if err == nil {
		t.Fatal("expected error when -t is missing")
	}
}

func TestParse_ValidTimestamp(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		timestamp int
		timezone  int
		format    string
	}{
		{
			name:      "timestamp only",
			args:      []string{"xtime-conv", "-t", "1710000000"},
			timestamp: 1710000000,
			timezone:  0,
			format:    DefaultFormat,
		},
		{
			name:      "with timezone",
			args:      []string{"xtime-conv", "-t", "1710000000", "-tz", "3"},
			timestamp: 1710000000,
			timezone:  3,
			format:    DefaultFormat,
		},
		{
			name:      "custom format",
			args:      []string{"xtime-conv", "-t", "1710000000", "-f", "2006-01-02"},
			timestamp: 1710000000,
			timezone:  0,
			format:    "2006-01-02",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetFlags(tt.args)

			cfg, err := Parse()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if cfg.Timestamp != tt.timestamp {
				t.Errorf("Timestamp = %d, want %d", cfg.Timestamp, tt.timestamp)
			}
			if cfg.Timezone != tt.timezone {
				t.Errorf("Timezone = %d, want %d", cfg.Timezone, tt.timezone)
			}
			if cfg.Format != tt.format {
				t.Errorf("Format = %q, want %q", cfg.Format, tt.format)
			}
		})
	}
}

func TestParse_ShowVersion(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{name: "short flag", args: []string{"xtime-conv", "-v"}},
		{name: "long flag", args: []string{"xtime-conv", "--version"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetFlags(tt.args)

			cfg, err := Parse()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !cfg.ShowVersion {
				t.Error("ShowVersion = false, want true")
			}
		})
	}
}
