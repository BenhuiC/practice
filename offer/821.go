package offer

func shortestToChar(s string, c byte) []int {
	n := len(s)
	res := make([]int, n)
	idx := -n
	for i, v := range s {
		if byte(v) == c {
			idx = i
		}
		res[i] = i - idx
	}

	idx = n * 2
	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			idx = i
		}
		res[i] = min(res[i], idx-i)
	}

	return res
}
