package partice

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) < 1 || len(obstacleGrid[0]) < 1 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	mp := make([][]int, m)
	for i := 0; i < m; i++ {
		mp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		mp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		mp[0][j] = 1
	}
	for x := 1; x < m; x++ {
		for y := 1; y < n; y++ {
			if obstacleGrid[x][y] == 1 {
				mp[x][y] = 0
				continue
			}
			if obstacleGrid[x-1][y] == 1 && obstacleGrid[x][y-1] == 1 {
				mp[x][y] = 0
			} else if obstacleGrid[x-1][y] == 1 {
				mp[x][y] = mp[x][y-1]
			} else if obstacleGrid[x][y-1] == 1 {
				mp[x][y] = mp[x-1][y]
			} else {
				mp[x][y] = mp[x-1][y] + mp[x][y-1]
			}
		}
	}
	return mp[m-1][n-1]
}
