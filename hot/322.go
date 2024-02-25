package hot

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
		for _, c := range coins {
			if i >= c && dp[i-c] != math.MaxInt {
				dp[i] = min(dp[i-c]+1, dp[i])
			}
		}
	}

	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
