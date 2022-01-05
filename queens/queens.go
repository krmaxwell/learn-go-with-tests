package queens

import (
	"fmt"
	"math"
)

type Board [][]bool

func CreateBoard(order int) Board {
	board := make([][]bool, order)
	for i := 0; i < order; i++ {
		board[i] = make([]bool, order)
	}
	return board
}

// solution based on https://medium.com/@egemenokte/solving-the-8-queens-puzzle-recursively-with-python-6440078b68ad

func PossibleLocation(b Board, rowCandidate, colCandidate int) bool {

	// test the actual location
	if b[rowCandidate][colCandidate] {
		return false
	}

	// test every row on the candidate column
	for _, row := range b {
		if row[colCandidate] {
			return false
		}
	}

	// test every column on the candidate row
	for column := range b[rowCandidate] {
		if b[rowCandidate][column] {
			return false
		}
	}

	// go through the whole board
	for row := range b {
		for column := range b[row] {
			if b[row][column] && math.Abs(float64(row-rowCandidate)) == math.Abs(float64(column-colCandidate)) {
				return false
			}
		}
	}
	return true
}

func CountQueens(b Board) int {
	count := 0
	for row := range b {
		for column := range b[row] {
			if b[row][column] {
				count++
			}
		}
	}
	return count
}

func SolveQueens(b Board) Board {
	for row := range b {
		for column := range b[row] {
			if PossibleLocation(b, row, column) {
				b[row][column] = true
				SolveQueens(b) // recursively solve
				if CountQueens(b) == len(b) {
					// we've placed all our queens
					return b
				}
				// this means we couldn't find a solution including this location
				b[row][column] = false
			}
		}
	}
	return Board{}
}

func PrintBoard(b Board) {
	for row := range b {
		for column := range b[row] {
			if b[row][column] {
				fmt.Print("Q")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
