package partice

import "fmt"

/*
每一天分两种情况
1.今天结束时手上有股票。2.今天结束是手上没有股票
dp[i][0]表示第i天手上没有股票时的最大受益
dp[i][1]表示第i天手上有股票时最大受益
今天手上有股票分为两种情况（1.昨天就有股票，今天没有卖。2.昨天没有股票，今天买入）
今天手上没有股票也分为两种情况（1.昨天没有股票，今天没有买入。2.昨天有股票，今天卖出）
*/

func maxProfit(prices []int) int {
	res122 := 0
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], prices[i]+dp[i-1][1])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	fmt.Println(dp)
	res122 = dp[len(prices)-1][0]
	return res122
}
