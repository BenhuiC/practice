package offer

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	locS := make([]int, 128)
	locT := make([]int, 128)
	n := len(s)
	for i := 0; i < n; i++ {
		if locS[s[i]] != locT[t[i]] {
			return false
		}
		locS[s[i]] = i + 1
		locT[t[i]] = i + 1
	}
	return true
}
