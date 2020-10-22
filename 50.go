package partice

import "math"

func myPow(x float64, n int) float64 {
	var result float64
	result = 1
	t := math.Abs(float64(n))
	for ; t > 0; t /= 2 {
		if int(t)%2 == 1 {
			result *= x
		}
		x *= x
	}
	if n < 0 {
		return 1 / result
	} else {
		return result
	}
}
