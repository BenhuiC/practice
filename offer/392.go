package offer

func isSubsequence(s string, t string) bool {
	var idx1, idx2 int
	l1, l2 := len(s), len(t)
	for ; idx1 < l1 && idx2 < l2; idx2++ {
		if s[idx1] == t[idx2] {
			idx1++
		}
	}

	return idx1 >= l1
}
