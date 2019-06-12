package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"bufio"
	"io"
	"log"
	"strconv"
)

func Printboard(board [][]int) {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if board[x][y] == 0 {
				fmt.Print("_")
			} else {
				fmt.Print(board[x][y])
			}
			if y == 2 || y == 5 {
				fmt.Print("|")
			} else if y == 8 {
				fmt.Print("\n")
			} else {
				fmt.Print(" ")
			}
		}
		if x == 2 || x == 5 {
			fmt.Println("-----+-----+-----")
		}
	}
}

func Reader(file string) [][]int {
	csvFile, _ := os.Open(file)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var board [][]int
	for x := 0; x < 9; x++ {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		var l []int
		board = append(board, l)
		var value int
		for y := 0; y < 9; y++ {
			value, _ = strconv.Atoi(line[y])
			board[x] = append(board[x], value)
		}
	}
	return board
}