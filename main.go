package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var file string
	flag.StringVar(&file, "file", "", "Board file")
	flag.Parse()

	if file == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	board := Reader(file)

	Printboard(board)
	fmt.Println()
	Solver(board)
}
