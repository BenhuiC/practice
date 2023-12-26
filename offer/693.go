package offer

func hasAlternatingBits(n int) bool {
	w := n & 1
	for n != 0 {
		t := n & 1
		if t^w == 1 {
			return false
		}
		n = n >> 1
		if w == 0 {
			w = 1
		} else {
			w = 0
		}
	}
	return true
}
