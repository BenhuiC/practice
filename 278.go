package partice

func firstBadVersion(n int) int {
	low, high := 1, n
	for low < high {
		t := (low + high) / 2
		if !isBadVersion(t-1) && isBadVersion(t) {
			return t
		} else if !isBadVersion(t) {
			low = t + 1
		} else {
			high = t
		}
	}
	return low
}

func isBadVersion(int) bool {
	return true
}
