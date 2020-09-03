package partice

func fib(N int) int {
	if N < 0 {
		return 0
	}
	if N <= 1 {
		return N
	}
	return fib(N-1) + fib(N-2)
}
