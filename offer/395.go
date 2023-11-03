package offer

import "strings"

func longestSubstring(s string, k int) int {
	m := make(map[int32]int)
	for _, v := range s {
		m[v-'a']++
	}
	var split byte
	for ch, v := range m {
		if v < k {
			split = byte(ch + 'a')
			break
		}
	}
	if split == 0 {
		return len(s)
	}
	res := 0
	for _, sub := range strings.Split(s, string(split)) {
		if subLen := longestSubstring(sub, k); subLen > res {
			res = subLen
		}
	}
	return res
}
