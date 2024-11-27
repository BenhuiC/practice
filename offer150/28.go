package offer150

func strStr(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	if m < n {
		return -1
	}
	for i := 0; i <= m-n; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}
