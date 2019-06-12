package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/mamoit/go-sudoku/pkg/solver"
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

	solver.Printboard(board)
	fmt.Println()
	solver.Solver(board)
}
