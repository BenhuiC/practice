package partice

func maxProfit714(prices []int, fee int) int {
	n := len(prices)
	sell, buy := 0, -prices[0]
	for i := 1; i < n; i++ {
		sell = Max(sell, buy+prices[i]-fee)
		buy = Max(buy, sell-prices[i])
	}
	return sell
}
