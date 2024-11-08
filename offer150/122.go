package offer150

func maxProfit122(prices []int) int {
	dp := [2][]int{make([]int, len(prices)), make([]int, len(prices))}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[0][i] = max(dp[0][i-1], dp[1][i-1]-prices[i])
		dp[1][i] = max(dp[1][i-1], prices[i]+dp[0][i-1])
	}

	return dp[1][len(prices)-1]
}
