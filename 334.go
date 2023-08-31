package partice

import "math"

func increasingTriplet(nums []int) bool {
	res334 := false
	l, n := math.MaxInt32, math.MaxInt32
	for _, v := range nums {
		if v <= l {
			l = v
			continue
		}
		if v <= n {
			n = v
			continue
		}
		res334 = true
		break
	}
	return res334
}
