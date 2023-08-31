package partice

func findAnagrams(s string, p string) []int {
	n := len(p)
	mp := make(map[byte]int)
	ms := make(map[byte]int)
	lastMatch := false
	for i := 0; i < n; i++ {
		ms[s[i]]++
		mp[p[i]]++
	}
	res := make([]int, 0)
	if match(ms, mp) {
		lastMatch = true
		res = append(res, 0)
	}
	for i := n; i < len(s); i++ {
		if s[i] == s[i-n] {
			if lastMatch {
				res = append(res, i-n+1)
			}
			continue
		}
		ms[s[i]]++
		ms[s[i-n]]--
		if lastMatch {
			lastMatch = false
			continue
		}
		if match(ms, mp) {
			lastMatch = true
			res = append(res, i-n+1)
		}
	}

	return res
}

func match(m1, m2 map[byte]int) bool {
	for k := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}
	return true
}
