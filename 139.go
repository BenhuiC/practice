package partice

func wordBreak(s string, wordDict []string) bool {
	mp := make(map[string]bool)
	for _, v := range wordDict {
		mp[v] = true
	}
	if mp[s] {
		return true
	}
	return isIn(s, mp)
}

func isIn(s string, mp map[string]bool) bool {
	if mp[s] {
		return true
	}
	for i := 1; i < len(s); i++ {
		if isIn(s[:i], mp) && isIn(s[i:], mp) {
			return true
		}
	}
	return false
}

func wordBreak2(s string, wordDict []string) bool {
	mp := make(map[string]bool)
	for _, v := range wordDict {
		mp[v] = true
	}
	if mp[s] {
		return true
	}
	result := make([]bool, len(s)+1)
	result[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if result[j] && mp[s[j:i]] {
				result[i] = true
				break
			}
		}
	}
	return result[len(s)]
}
