package partice

import "math"

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if -2147483648 == dividend && -1 == divisor {
		return 2147483647
	}
	var result int
	negative := dividend^divisor < 0
	d1, d2 := int(math.Abs(float64(dividend))), int(math.Abs(float64(divisor)))
	for i := 31; i >= 0; i-- {
		if d1>>i >= d2 {
			result += 1 << i
			d1 -= d2 << i
		}
	}
	if negative {
		return -result
	}
	return result
}
