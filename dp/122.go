package dp

func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}

	return res
}

/*
dp[i][0] 表示为第i天卖出股票或不操作的最大收益
dp[i][1] 表示为第i天买入股票的最大收益
*/
func maxProfit2(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	dp := make([][2]int, n)

	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i]) // 上一天买入或没有操作，这一天可以卖出
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]-prices[i]) // 上一天卖出或没有操作，这一天可以买入
	}

	return dp[n-1][0]
}

// 空间优化
func maxProfit3(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	sell := 0
	buy := -prices[0]
	for i := 1; i < n; i++ {
		// 上一天买入或没有操作，这一天可以卖出
		// 上一天卖出或没有操作，这一天可以买入
		sell, buy = Max(sell, buy+prices[i]), Max(buy, sell-prices[i])
	}

	return sell
}
