package array

func camelMatch(queries []string, pattern string) []bool {
	match := func(q, p string) bool {
		idx1, idx2 := 0, 0
		for idx1 < len(q) && idx2 < len(p) {
			if q[idx1] == p[idx2] {
				idx1++
				idx2++
			} else if q[idx1] <= 'z' && q[idx1] >= 'a' {
				idx1++
			} else {
				return false
			}
		}
		if idx2 < len(p) {
			return false
		}
		for idx1 < len(q) {
			if q[idx1] <= 'Z' && q[idx1] >= 'A' {
				return false
			}
			idx1++
		}
		return true
	}
	res := make([]bool, len(queries))
	for i, q := range queries {
		res[i] = match(q, pattern)
	}

	return res
}
