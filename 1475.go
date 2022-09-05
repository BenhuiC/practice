package partice

func finalPrices(prices []int) []int {
	stk := make([]int, 0)
	for i, v := range prices {
		if len(stk) == 0 {
			stk = append(stk, i)
			continue
		}
		for len(stk) > 0 && prices[stk[len(stk)-1]] >= v {
			t := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			prices[t] -= v
		}
		stk = append(stk, i)
	}

	return prices
}
