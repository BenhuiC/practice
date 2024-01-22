package offer

func minimumLengthEncoding(words []string) int {
	res := 0
	m := make(map[string]struct{})
	for _, w := range words {
		m[w] = struct{}{}
	}

	for _, w := range words {
		for i := 1; i < len(w); i++ {
			if _, ok := m[w[i:]]; ok {
				delete(m, w[i:])
			}
		}
	}

	for k := range m {
		res += len(k) + 1
	}

	return res
}
