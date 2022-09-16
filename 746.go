package partice

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	for i := 0; i < len(cost); i++ {
		if i < 2 {
			dp[i] = cost[i]
			continue
		}
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return min(dp[len(cost)-1], dp[len(cost)-2])
}
