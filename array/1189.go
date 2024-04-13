package array

func maxNumberOfBalloons(text string) int {
	// balloon
	m := map[int32]int{'b': 0, 'a': 0, 'l': 0, 'o': 0, 'n': 0}
	for _, c := range text {
		if _, ok := m[c]; ok {
			m[c]++
		}
	}
	res := len(text) / 7
	for k, v := range m {
		if k == 'l' || k == 'o' {
			res = min(res, v/2)
		} else {
			res = min(res, v)
		}
	}

	return res
}
