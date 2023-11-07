package offer

func countBattleships(board [][]byte) int {
	res := 0
	direction := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(board), len(board[0])
	queue := make([][2]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' || board[i][j] == ' ' {
				continue
			}
			res++
			queue = queue[:0]
			queue = append(queue, [2]int{i, j})
			for len(queue) != 0 {
				item := queue[0]
				queue = queue[1:]
				board[item[0]][item[1]] = ' '
				for _, d := range direction {
					x, y := item[0]+d[0], item[1]+d[1]
					if x < 0 || x >= m || y < 0 || y >= n {
						continue
					}
					if board[x][y] == 'X' {
						queue = append(queue, [2]int{x, y})
					}
				}
			}
		}
	}
	return res
}
