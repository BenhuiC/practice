package partice

func maxProfit121(prices []int) int {
	res121 := 0
	m := prices[0]
	for _, v := range prices {
		m = min(m, v)
		res121 = max(res121, v-m)
	}

	return res121
}
