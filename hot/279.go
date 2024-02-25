package hot

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = n
		for j := 1; j <= int(math.Sqrt(float64(i))); j++ {
			dp[i] = min(dp[i], 1+dp[i-j*j])
		}
	}

	return dp[n]
}
