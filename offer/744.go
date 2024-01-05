package offer

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	if n == 0 {
		return ' '
	}
	if target < letters[0] || target >= letters[n-1] {
		return letters[0]
	}
	l, r := 0, n-1
	for l < r {
		mid := l + (r-l)>>1
		if letters[mid] > target && letters[mid-1] <= target {
			return letters[mid]
		} else if letters[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return letters[l]
}
