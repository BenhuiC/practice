package dp

/*
dp[i][j]表示为从坐标[i,j]到达终点所需要的最小点数
dp[i][j]=Max(1,Min(dp[i+1][j],dp[i][j+1]))
*/
func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	if m == 0 {
		return 0
	}
	n := len(dungeon[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[m-1][n-1] = Max(1, 1-dungeon[m-1][n-1])
	for i := m - 2; i >= 0; i-- {
		dp[i][n-1] = Max(1, dp[i+1][n-1]-dungeon[i][n-1])
	}
	for j := n - 2; j >= 0; j-- {
		dp[m-1][j] = Max(1, dp[m-1][j+1]-dungeon[m-1][j])
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			minn := Min(dp[i][j+1], dp[i+1][j])
			dp[i][j] = Max(1, minn-dungeon[i][j])
		}
	}

	return dp[0][0]
}
