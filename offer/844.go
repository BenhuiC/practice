package offer

func backspaceCompare(s string, t string) bool {
	s1 := []byte(s)
	s2 := []byte(t)
	idx1, idx2 := 0, 0
	for _, c := range s {
		if c == '#' {
			if idx1 > 0 {
				idx1--
			}
			continue
		}
		s1[idx1] = byte(c)
		idx1++
	}

	for _, c := range t {
		if c == '#' {
			if idx2 > 0 {
				idx2--
			}
			continue
		}
		s2[idx2] = byte(c)
		idx2++
	}

	return string(s1[:idx1]) == string(s2[:idx2])
}
