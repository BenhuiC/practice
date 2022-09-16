package offer

import "math"

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	res14 := 1
	t, y := n/3, n%3
	if y == 1 {
		res14 = 4
		t--
	} else if y != 0 {
		res14 = y
	}
	res14 = res14 * int(math.Pow(3.0, float64(t)))
	return res14
}
