package clever

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	count := 0
	for n != 0 {
		if n&1 == 1 {
			count++
		}
		n = n >> 1
		if count > 1 {
			return false
		}
	}
	return count == 1
}
