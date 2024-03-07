package dp

import "math"

func maxProfit188(k int, prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	k = min(k, n/2)
	/*
		buy[i][j] 表示在price[0:i]的价格中交易了j次时手中还持有股票的最大收益
		sell[i][j] 表示在price[0:i]的价格中交易了j次时手中不持有股票的最大收益

		buy[i][j]=max(buy[i-1][j],sell[i-1][j]-prices[i]) 当天购买股票和不购买股票的最大值
		sell[i][j]=max(sell[i-1][j],buy[i-1][j-1]+prices[i]) 当天出售股票和不出售股票的最大值
	*/
	buy, sell := make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		buy[i] = make([]int, k+1)
		sell[i] = make([]int, k+1)
	}
	buy[0][0] = -prices[0]
	for i := 1; i <= k; i++ {
		buy[0][i] = math.MinInt / 2
		sell[0][i] = math.MinInt / 2
	}

	for i := 1; i < n; i++ {
		buy[i][0] = max(buy[i-1][0], -prices[i])
		for j := 1; j <= k; j++ {
			buy[i][j] = max(buy[i-1][j], sell[i-1][j]-prices[i])
			sell[i][j] = max(sell[i-1][j], buy[i-1][j-1]+prices[i])
		}
	}

	res := max(sell[n-1]...)

	return res
}
