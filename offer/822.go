package offer

import "math"

func flipgame(fronts []int, backs []int) int {
	same := make(map[int]bool)
	for i := range fronts {
		if fronts[i] == backs[i] {
			same[fronts[i]] = true
		}
	}

	res := math.MaxInt
	for _, v := range fronts {
		if v < res && !same[v] {
			res = v
		}
	}

	for _, v := range backs {
		if v < res && !same[v] {
			res = v
		}
	}
	if res == math.MaxInt {
		return 0
	}
	return res
}
