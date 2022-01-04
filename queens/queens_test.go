package queens

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueens(t *testing.T) {
	order := 3

	board := PlaceQueens(order)
	want := 3

	assert.Len(t, board, want)
	assert.Equal(t, CountQueensInSlice(t, board[0]), 1)
}

func CountQueensInSlice(t *testing.T, vector []bool) int {
	t.Helper()

	count := 0
	for _, v := range vector {
		if v {
			count++
		}
	}

	return count
}
