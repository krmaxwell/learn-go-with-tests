package queens

type Board [][]bool

func CreateBoard(order int) Board {
	board := make([][]bool, order)
	for i := 0; i < order; i++ {
		board[i] = make([]bool, order)
	}
	return board
}
