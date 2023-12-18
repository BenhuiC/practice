package offer

import "math"

func judgeSquareSum(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		tmp := left*left + right*right
		if tmp == c {
			return true
		} else if tmp > c {
			right--
		} else {
			left++
		}
	}
	return false
}
