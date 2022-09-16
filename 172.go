package partice

func trailingZeroes(n int) int {
	var res172 int
	var num2, num5 int
	for i := 2; i <= n; i++ {
		for t := i; t%2 == 0; t = t / 2 {
			num2++
		}
		for t := i; t%5 == 0; t = t / 5 {
			num5++
		}
	}
	res172 = min(num2, num5)
	return res172
}
