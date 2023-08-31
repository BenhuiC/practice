package partice

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	var res204 int
	ary := make([]bool, n)
	for i := range ary {
		ary[i] = true
	}
	for i := 2; i < n; i++ {
		if ary[i] {
			res204++
			for j := 2 * i; j < n; j += i {
				ary[j] = false
			}
		}
	}
	return res204
}
