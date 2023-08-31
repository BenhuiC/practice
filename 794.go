package partice

// x数量大于等于o且最多大1
// 如果x胜，则o数量为x-1
// 如果o胜，则x，o数量相等
// x，o只能一方胜利

func validTicTacToe(board []string) bool {
	xNum := 0
	oNum := 0
	for _, ar := range board {
		for _, v := range ar {
			if v == 'X' {
				xNum++
			} else if v == 'O' {
				oNum++
			}
		}
	}
	if xNum-oNum > 1 || xNum-oNum < 0 {
		return false
	}
	xWin := win('X', board)
	oWin := win('O', board)
	if xWin {
		return !oWin && xNum == oNum+1
	}
	if win('O', board) {
		return xNum == oNum
	}
	return true
}

func win(ch byte, board []string) bool {
	for i := 0; i < 3; i++ {
		nux := 0
		nuy := 0
		for j := 0; j < 3; j++ {
			if board[i][j] == ch {
				nux++
			}
			if board[j][i] == ch {
				nuy++
			}
		}
		if nux == 3 || nuy == 3 {
			return true
		}
	}
	x, y := 0, 0
	for i := 0; i < 3; i++ {
		if board[i][i] == ch {
			x++
		}
		if board[2-i][i] == ch {
			y++
		}
	}
	if x == 3 || y == 3 {
		return true
	}
	return false
}
