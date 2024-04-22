package array

import "sort"

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	low, hi := 0, len(tokens)-1
	var point, res int
	for low <= hi && (power >= tokens[low] || point > 0) {
		for low <= hi && power >= tokens[low] {
			power -= tokens[low]
			low++
			point++
		}
		res = max(res, point)
		if low <= hi && point > 0 {
			power += tokens[hi]
			hi--
			point--
		}
	}

	return res
}
