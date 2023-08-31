package partice

func longestPalindrome409(s string) int {
	res409 := 0
	m := make(map[uint8]int)
	for i := range s {
		m[s[i]]++
	}
	hasO := false
	for k, v := range m {
		if v%2 == 1 {
			hasO = true
			m[k]--
		}
	}
	for _, v := range m {
		res409 += v
	}
	if hasO {
		res409++
	}

	return res409
}
