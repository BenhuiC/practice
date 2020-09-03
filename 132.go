package partice

func minCut(s string) int {
	if isHui(s) {
		return 0
	}
	result := 1 + minCut(s[:1]) + minCut(s[1:])
	for i := 2; i < len(s); i++ {
		v := 1 + minCut(s[:i]) + minCut(s[i:])
		if v < result {
			result = v
		}
	}
	return result
}

func isHui(s string) bool {
	if len(s) <= 1 {
		return true
	}
	for i, j := 0, len(s)-1; i <= j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
