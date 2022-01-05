package queens

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardCreation(t *testing.T) {
	order := 3

	board := CreateBoard(order)
	want := 3

	assert.Len(t, board, want)
	assert.Len(t, board[0], want)
	assert.IsType(t, Board{}, board)
	assert.False(t, board[0][0])
}

func TestValidation(t *testing.T) {
	validBoard := Board{
		{true, false, false},
		{false, false, true},
		{false, true, false},
	}
	assert.True(t, ValidateBoard(validBoard))

	allFalseBoard := Board{
		{false, false, false},
		{false, false, false},
		{false, false, false},
	}
	assert.False(t, ValidateBoard(allFalseBoard))

	allTrueBoard := Board{
		{true, true, true},
		{true, true, true},
		{true, true, true},
	}
	assert.False(t, ValidateBoard(allTrueBoard))

	firstColumnBoard := Board{
		{true, false, false},
		{true, false, false},
		{true, false, false},
	}
	assert.False(t, ValidateBoard(firstColumnBoard))

	firstRowBoard := Board{
		{true, true, true},
		{false, false, false},
		{false, false, false},
	}
	assert.False(t, ValidateBoard(firstRowBoard))

}
