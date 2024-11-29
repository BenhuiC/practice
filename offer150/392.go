package offer150

func isSubsequence(s string, t string) bool {
	idx1, idx2 := 0, 0
	for idx1 < len(s) && idx2 < len(t) {
		if s[idx1] == t[idx2] {
			idx1++
			idx2++
		} else {
			idx2++
		}
	}
	return idx1 == len(s)
}
