package main

import (
	"fmt"
	"os"
	"xtime-conv/internal/config"
	"xtime-conv/internal/parser"
	"xtime-conv/internal/timeutil"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	tm := parser.ParseFlexibleUnix(int64(cfg.Timestamp)).UTC()

	fmt.Printf("UTC:\t%v\n", timeutil.InLocalTimastamp(tm, 0).Format(cfg.Format))

	if cfg.Timezone != 0 {
		fmt.Printf("Local:\t%v\n", timeutil.InLocalTimastamp(tm, cfg.Timezone).Format(cfg.Format))
	}
}
