package dp

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	dp := make([]int, len(grid[0]))
	dp[0] = grid[0][0]
	for i := 1; i < len(dp); i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}
	for i := 1; i < len(grid); i++ {
		dp[0] += grid[i][0]
		for j := 1; j < len(grid[0]); j++ {
			dp[j] = Min(dp[j-1], dp[j]) + grid[i][j]
		}
	}
	return dp[len(grid[0])-1]
}
