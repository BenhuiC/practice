package array

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	left, right := 0, 0
	m := make(map[byte]struct{})
	m[s[left]] = struct{}{}
	res := 1
	for right < len(s)-1 {
		if _, ok := m[s[right+1]]; !ok {
			right++
			if right-left+1 > res {
				res = right - left + 1
			}
			m[s[right]] = struct{}{}
		} else {
			delete(m, s[left])
			left++
		}
	}

	return res
}
