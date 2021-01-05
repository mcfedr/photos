package main

import (
	"os"
	"path/filepath"
)
import "flag"

func main() {
	var batch = flag.Int("batch", 1000, "Size of the batches to report")
	var help = flag.Bool("help", false, "show usage")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(1)
	}

	dir := flag.Arg(0)
	if dir == "" {
		println("You should specify path to scan")
		os.Exit(1)
	}
	c := 0
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		c++
		if c % *batch == 0 {
			println(c, "scanning:", path)
		}

		return nil
	})
	if err != nil {
		println("Error scanning dir:", err)
		os.Exit(1)
	}

	println("scanned:", c, "files and dirs")
}
