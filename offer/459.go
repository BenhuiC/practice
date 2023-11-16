package offer

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	for i := 1; i <= n/2; i++ {
		if n%i != 0 || n/i <= 1 {
			continue
		}
		sub := s[:i]
		match := true
		for j := 0; j < n; j += i {
			if sub != s[j:j+i] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}
