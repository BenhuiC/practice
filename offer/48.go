package offer

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	res := 1
	m := map[uint8]bool{}
	m[s[0]] = true
	l, r := 0, 0
	for r < len(s)-1 {
		if !m[s[r+1]] {
			m[s[r+1]] = true
			r++
			res = max(res, len(m))
			continue
		}
		delete(m, s[l])
		l++
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
