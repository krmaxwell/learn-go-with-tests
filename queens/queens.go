package queens

type Board [][]bool

func CreateBoard(order int) Board {
	board := make([][]bool, order)
	for i := 0; i < order; i++ {
		board[i] = make([]bool, order)
	}
	return board
}

func ValidateBoard(b Board) bool {
	// returns true if board has one true value on each row, column, and diagonal

	for _, row := range b {
		if !validateSlice(row) {
			return false
		}
	}

	for col := range b[0] {
		colSlice := make([]bool, len(b[0]))
		for row := range b {
			colSlice = append(colSlice, b[row][col])
		}
		if !validateSlice(colSlice) {
			return false
		}
	}

	return true
}

func validateSlice(s []bool) bool {
	count := 0
	for _, cell := range s {
		if cell {
			count++
		}
	}
	return count == 1
}
