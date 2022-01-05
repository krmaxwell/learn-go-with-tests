package queens

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardCreation(t *testing.T) {
	order := 8

	board := CreateBoard(order)
	want := 8

	assert.Len(t, board, want)
	assert.Len(t, board[0], want)
	assert.IsType(t, Board{}, board)
	assert.False(t, board[0][0])
}

func TestPossibleLocation(t *testing.T) {
	board := CreateBoard(8)
	board[1][1] = true
	assert.False(t, PossibleLocation(board, 0, 1)) // can't put on the same column
	assert.False(t, PossibleLocation(board, 1, 0)) // can't put on the same row
	assert.False(t, PossibleLocation(board, 4, 4)) // can't put on the same diagonal
	assert.True(t, PossibleLocation(board, 2, 3))
}

func TestCountQueens(t *testing.T) {
	board := CreateBoard(8)
	assert.Equal(t, 0, CountQueens(board))
	board[1][1] = true
	assert.Equal(t, 1, CountQueens(board))
	board[3][7] = true
	assert.Equal(t, 2, CountQueens(board))
}

func TestSolve(t *testing.T) {
	board := CreateBoard(8)
	solution := SolveQueens(board)
	assert.Equal(t, 8, CountQueens(solution))
	PrintBoard(solution)
}
