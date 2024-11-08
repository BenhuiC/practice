package offer150

func maxProfit(prices []int) int {
	res := 0
	m := prices[0]

	for i := 1; i < len(prices); i++ {
		m = min(m, prices[i])
		res = max(res, prices[i]-m)
	}
	return res
}
