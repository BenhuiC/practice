package partice

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	n := len(s1)
	diffCnt := 0
	lastDiffIdx := 0
	for i := 0; i < n; i++ {
		if s1[i] == s2[i] {
			continue
		}
		diffCnt++
		if diffCnt > 2 {
			return false
		} else if diffCnt == 1 {
			lastDiffIdx = i
		} else {
			if s1[lastDiffIdx] == s2[i] && s2[lastDiffIdx] == s1[i] {
				continue
			}
			return false
		}
	}
	if diffCnt == 1 || diffCnt > 2 {
		return false
	}
	return true
}
