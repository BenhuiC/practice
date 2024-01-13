package offer

func rotateString(s string, goal string) bool {
	if s == goal {
		return true
	}
	l1, l2 := len(s), len(goal)
	if l1 != l2 {
		return false
	}

	for i := 1; i < l1; i++ {
		if s[i:] == goal[:l1-i] && s[:i] == goal[l1-i:] {
			return true
		}
	}
	return false
}
