package hot

func lengthOfLongestSubstring(s string) int {
	exist := make(map[byte]int)
	l, r := 0, 0
	n := len(s)
	res := 0
	for r < n {
		exist[s[r]] += 1
		for l < r && exist[s[r]] > 1 {
			exist[s[l]]--
			l++
		}
		if r-l+1 > res {
			res = r - l + 1
		}
		r++
	}

	return res
}
