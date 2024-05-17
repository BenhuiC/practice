package array

func numRookCaptures(board [][]byte) int {
	m := len(board)
	if m == 0 {
		return 0
	}
	n := len(board[0])

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != 'R' {
				continue
			}
			// left
			left := j - 1
			for left >= 0 && board[i][left] == '.' {
				left--
			}
			if left >= 0 && board[i][left] == 'p' {
				res++
			}
			// right
			right := j + 1
			for right < n && board[i][right] == '.' {
				right++
			}
			if right < n && board[i][right] == 'p' {
				res++
			}

			// up
			up := i - 1
			for up >= 0 && board[up][j] == '.' {
				up--
			}
			if up >= 0 && board[up][j] == 'p' {
				res++
			}

			// down
			down := i + 1
			for down < m && board[down][j] == '.' {
				down++
			}
			if down < m && board[down][j] == 'p' {
				res++
			}
		}
	}

	return res
}
