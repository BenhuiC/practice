package partice

func coinChange(coins []int, amount int) int {
	MAX_INT := int(^uint(0) >> 1)
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		cost := MAX_INT
		for _, v := range coins {
			if i-v >= 0 && dp[i-v] != MAX_INT {
				cost = Min(cost, dp[i-v]+1)
			}
		}
		dp[i] = cost
	}
	if dp[amount] == MAX_INT {
		return -1
	}
	return dp[amount]
}
