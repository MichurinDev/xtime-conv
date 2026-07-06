package main

import (
	"fmt"
	"os"

	"github.com/MichurinDev/xtime-conv/internal/config"
	"github.com/MichurinDev/xtime-conv/internal/parser"
	"github.com/MichurinDev/xtime-conv/internal/timeutil"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if cfg.ShowVersion {
		fmt.Printf("xtime-conv %s (commit: %s, built: %s)\n", version, commit, date)
		os.Exit(0)
	}

	tm := parser.ParseFlexibleUnix(int64(cfg.Timestamp)).UTC()

	fmt.Printf("UTC:\t%v\n", timeutil.InLocalTimestamp(tm, 0).Format(cfg.Format))

	if cfg.Timezone != 0 {
		fmt.Printf("Local:\t%v\n", timeutil.InLocalTimestamp(tm, cfg.Timezone).Format(cfg.Format))
	}
}
