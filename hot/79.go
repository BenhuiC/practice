package hot

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	res := false
	var bk func(x, y, idx int)
	bk = func(x, y, idx int) {
		if res {
			return
		}
		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != word[idx] {
			return
		}
		if idx == len(word)-1 {
			res = true
			return
		}
		idx++
		t := board[x][y]
		board[x][y] = '0'
		bk(x, y+1, idx)
		bk(x, y-1, idx)
		bk(x-1, y, idx)
		bk(x+1, y, idx)
		board[x][y] = t
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				bk(i, j, 0)
			}
			if res {
				return true
			}
		}
	}

	return res
}
