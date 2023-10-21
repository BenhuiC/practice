package offer

import (
	"math"
)

func isUgly(n int) bool {
	t := int(math.Sqrt(float64(n)))
	can := make(map[int]bool)
	can[2] = true
	can[3] = true
	can[5] = true
	for i := 2; i <= t; i++ {
		if n%i == 0 && (!can[i] || !can[n/i]) {
			return false
		}
	}
	return true
}
