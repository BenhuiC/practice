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
				dp[i][j] = matrix[i][j] + Min(dp[i-1][j], dp[i-1][j+1])
			} else if j == n-1 {
				dp[i][j] = matrix[i][j] + Min(dp[i-1][j-1], dp[i-1][j])
			} else {
				dp[i][j] = matrix[i][j] + Min(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1])
			}
		}
	}

	for _, v := range dp[n-1] {
		res = Min(res, v)
	}
	return res
}

func minFallingPathSum2(matrix [][]int) int {
	n := len(matrix)
	dp := matrix[0]
	for i := 1; i < n; i++ {
		tmp := make([]int, n)
		for j := 0; j < n; j++ {
			if j == 0 {
				tmp[j] = matrix[i][j] + Min(dp[j], dp[j+1])
			} else if j == n-1 {
				tmp[j] = matrix[i][j] + Min(dp[j-1], dp[j])
			} else {
				tmp[j] = matrix[i][j] + Min(dp[j-1], dp[j], dp[j+1])
			}
		}
		dp = tmp
	}

	res := Min(dp...)
	return res
}
