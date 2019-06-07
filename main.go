package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
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

func main() {
	csvFile, _ := os.Open("board.csv")
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

	Printboard(board)
	fmt.Println()
	Solver(board)
}

func CheckValid(board [][]int, x int, y int) []int {
	var invalid []int

	for a := 0; a < 9; a++ {
		// Check column
		if board[a][y] != 0 {
			invalid = append(invalid, board[a][y])
		}
		// Check row
		if board[x][a] != 0 {
			invalid = append(invalid, board[x][a])
		}
	}

	// Check 3x3 square
	x_i := x / 3
	y_i := y / 3

	for i := x_i * 3; i < (x_i+1)*3; i++ {
		for j := y_i * 3; j < (y_i+1)*3; j++ {
			if board[i][j] != 0 {
				invalid = append(invalid, board[i][j])
			}
		}
	}

	sort.Ints(invalid)

	var valid []int
	i := 0

	for a := 1; a <= 9; a++ {
		for invalid[i] < a {
			if len(invalid)-1 == i {
				break
			}
			i++
		}
		if invalid[i] != a {
			valid = append(valid, a)
		}
	}

	return valid
}

func Solver(board [][]int) bool {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if board[x][y] == 0 {
				valid := CheckValid(board, x, y)
				for i := 0; i < len(valid); i++ {
					board[x][y] = valid[i]
					if Solver(board) == true {
						return true
					}
				}
				board[x][y] = 0
				return false
			}
		}
	}
	Printboard(board)

	return true
}
