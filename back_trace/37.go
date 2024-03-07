package back_trace

func solveSudoku(board [][]byte) {
	row := make([][]bool, 9)
	col := make([][]bool, 9)
	grid := make([][]bool, 9)
	space := make([][2]int, 0)
	for i := 0; i < 9; i++ {
		row[i] = make([]bool, 9)
		col[i] = make([]bool, 9)
		grid[i] = make([]bool, 9)
	}
	gridIdx := func(x, y int) int {
		r := x / 3
		c := y / 3
		return 3*r + c
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				row[i][board[i][j]-'1'] = true
				col[j][board[i][j]-'1'] = true
				grid[gridIdx(i, j)][board[i][j]-'1'] = true
			} else {
				space = append(space, [2]int{i, j})
			}
		}
	}

	var bk func(pos int) bool
	bk = func(pos int) bool {
		if pos == len(space) {
			return true
		}
		for i := 1; i <= 9; i++ {
			r, c := space[pos][0], space[pos][1]
			if row[r][i-1] || col[c][i-1] || grid[gridIdx(r, c)][i-1] {
				continue
			}
			row[r][i-1] = true
			col[c][i-1] = true
			grid[gridIdx(r, c)][i-1] = true
			board[r][c] = byte('0' + i)
			if bk(pos + 1) {
				return true
			}
			row[r][i-1] = false
			col[c][i-1] = false
			grid[gridIdx(r, c)][i-1] = false
			board[r][c] = '0'
		}

		return false
	}
	bk(0)
}
