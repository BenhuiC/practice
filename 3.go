package partice

func lengthOfLongestSubstring(s string) int {
	var max int

	for i := 0; i < len(s); i++ {
		m := make(map[int32]bool)
		n := 0
		for _, a := range s[i:] {
			if _, ok := m[a]; !ok {
				n++
				m[a] = true
			} else {
				break
			}
		}
		if n > max {
			max = n
		}
	}

	return max
}
func lengthOfLongestSubstring2(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	max := 1
	start, end := 0, 0
	m := make([]int, 256)
	for end < len(s) {
		if end < len(s) && m[s[end]] == 0 {
			m[s[end]] = 1
			end++
		} else {
			m[s[start]] = 0
			start++
		}
		n := end - start
		if n > max {
			max = n
		}
	}
	return max
}

func lengthOfLongestSubstring3(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	max := 1
	start, end := 0, 0
	m := make(map[uint8]bool)
	for end < len(s) {
		if yes, ok := m[s[end]]; !ok || !yes {
			m[s[end]] = true
			end++
		} else {
			m[s[start]] = false
			start++
		}
		n := end - start
		if n > max {
			max = n
		}
	}
	return max
}
