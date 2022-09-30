package interview

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	for i := 1; i < len(s1); i++ {
		if s2 == s1[i:]+s1[:i] {
			return true
		}
	}

	return false
}
