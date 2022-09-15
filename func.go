package partice

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func absFloat64(f float64) float64 {
	if f < 0 {
		return -f
	}
	return f
}
