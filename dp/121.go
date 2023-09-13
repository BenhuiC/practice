package dp

func maxProfit121(prices []int) int {
	res := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		min = Min(prices[i], min)
		res = Max(res, prices[i]-min)
	}
	return res
}
