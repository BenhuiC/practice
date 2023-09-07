package dp

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, v := range wordDict {
		wordMap[v] = true
	}
	if wordMap[s] {
		return true
	}
	slen := len(s)
	dp := make([]bool, slen+1)
	dp[0] = true
	for r := 1; r <= slen; r++ {
		for l := 0; l < r; l++ {
			if dp[l] && wordMap[s[l:r]] {
				dp[r] = true
				break
			}
		}
	}
	return dp[slen]
}

// 这个方法逻辑上是对的，但是leetcode会超时
func wordBreak2(s string, wordDict []string) bool {
	slen := len(s)
	wlen := len(wordDict)
	if slen == 0 || wlen == 0 {
		return false
	}

	var fn func(l int) bool
	fn = func(l int) bool {
		if l >= slen {
			return true
		}
		r := false
		for _, v := range wordDict {
			wl := len(v)
			if l+wl > slen {
				continue
			}
			if s[l:l+wl] == v && fn(l+wl) {
				r = true
				break
			}
		}
		return r
	}

	return fn(0)
}
