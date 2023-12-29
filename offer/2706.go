package offer

import "math"

func buyChoco(prices []int, money int) int {
	first, second := math.MaxInt, math.MaxInt
	for _, v := range prices {
		if v < first {
			first, second = v, first
		} else if v < second {
			second = v
		}
	}

	diff := money - first - second
	if diff < 0 {
		return money
	}
	return diff
}
