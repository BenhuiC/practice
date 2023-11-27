package offer

import "math"

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	n := int(math.Sqrt(float64(num)))
	sum := 0
	for i := 1; i <= n; i++ {
		if num%i != 0 {
			continue
		}
		y := num / i
		sum += i
		if y != num {
			sum += y
		}
	}
	return sum == num
}
