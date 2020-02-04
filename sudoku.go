package main

import (
	"flag"
	"github.com/mamoit/go-sudoku/pkg/solver"
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

	board := solver.Reader(file)

	solver.Solver(board)
}
