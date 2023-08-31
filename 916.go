package partice

func wordSubsets(words1 []string, words2 []string) []string {
	n1 := len(words1)
	n2 := len(words2)
	words1Map := make([]map[int32]int, n1)
	for x := 0; x < n1; x++ {
		m := make(map[int32]int)
		for _, v := range words1[x] {
			m[v]++
		}
		words1Map[x] = m
	}
	w2m := make(map[int32]int)
	for x := 0; x < n2; x++ {
		m := make(map[int32]int)
		for _, v := range words2[x] {
			m[v]++
			if w2m[v] < m[v] {
				w2m[v] = m[v]
			}
		}

	}
	res916 := make([]string, 0)
	for i := 0; i < n1; i++ {
		if !contain(words1Map[i], w2m) {
			continue
		}

		res916 = append(res916, words1[i])

	}
	return res916
}

func contain(m1, m2 map[int32]int) bool {
	for k, v := range m2 {
		if m1[k] < v {
			return false
		}
	}
	return true
}
