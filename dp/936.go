package dp

import "math"

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	dp := make([][]int, n)
	res := math.MaxInt
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[i][j] = matrix[i][j]
				continue
			}
			if j == 0 {
				dp[i][j] = matrix[i][j] + Mins(dp[i-1][j], dp[i-1][j+1])
			} else if j == n-1 {
				dp[i][j] = matrix[i][j] + Mins(dp[i-1][j-1], dp[i-1][j])
			} else {
				dp[i][j] = matrix[i][j] + Mins(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1])
			}
		}
	}

	for _, v := range dp[n-1] {
		res = Mins(res, v)
	}
	return res
}
