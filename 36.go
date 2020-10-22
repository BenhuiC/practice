package partice

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if !isValidLine(board, i) || !isValidColumn(board, i) {
			return false
		}
	}
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if !isValidCube(board, 1+x*3, 1+y*3) {
				return false
			}
		}
	}
	return true
}

func isValidLine(board [][]byte, i int) bool {
	m := make(map[byte]bool, 0)
	ary := board[i]
	for j := 0; j < len(ary); j++ {
		if ary[j] == '.' {
			continue
		}
		if m[ary[j]] {
			return false
		}
		m[ary[j]] = true
	}
	return true
}

func isValidColumn(board [][]byte, j int) bool {
	m := make(map[byte]bool, 0)
	for i := 0; i < len(board); i++ {
		if board[i][j] == '.' {
			continue
		}
		if m[board[i][j]] {
			return false
		}
		m[board[i][j]] = true
	}
	return true
}

func isValidCube(board [][]byte, i, j int) bool {
	m := make(map[byte]bool, 0)
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if board[i+x][j+y] == '.' {
				continue
			}
			if m[board[i+x][j+y]] {
				return false
			}
			m[board[i+x][j+y]] = true
		}
	}
	return true
}
