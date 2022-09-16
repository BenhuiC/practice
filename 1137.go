package partice

func tribonacci(n int) int {
	ary := make([]int, n+3)
	ary[1] = 1
	ary[2] = 1
	var rec func(n int) int
	rec = func(n int) int {
		if n >= len(ary) || n < 0 {
			return 0
		}
		if ary[n] != 0 {
			return ary[n]
		}
		t := rec(n-1) + rec(n-2) + rec(n-3)
		ary[n] = t
		return t
	}
	return rec(n)
}
