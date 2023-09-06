package dp

func fib(n int) int {
	if n <= 1 {
		return n
	}
	last2, last1 := 0, 1
	for i := 2; i <= n; i++ {
		last2, last1 = last1, last1+last2
	}
	return last1
}
