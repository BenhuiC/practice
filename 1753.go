package partice

func maximumScore(a int, b int, c int) int {
	min, mid, max := a, b, c
	if mid < min {
		mid, min = min, mid
	}
	if max < min {
		min, max = max, min
	}
	if mid > max {
		max, mid = mid, max
	}

	if min+mid > max {
		return max + (min+mid-max)/2
	}
	return min + mid
}
