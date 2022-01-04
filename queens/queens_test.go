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
