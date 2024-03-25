package array

func isHappy(n int) bool {
	m := make(map[int]bool)
	var next int
	for {
		next = 0
		for n != 0 {
			v := n % 10
			next += v * v
			n /= 10
		}
		if next == 1 {
			return true
		}
		if m[next] {
			return false
		}
		m[next] = true
		n = next
	}
}
