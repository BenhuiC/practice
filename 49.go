package partice

func groupAnagrams(strs []string) [][]string {
	m := map[int32]int64{'a': 2, 'b': 3, 'c': 5, 'd': 7, 'e': 11, 'f': 13, 'g': 17, 'h': 19, 'i': 23, 'j': 29, 'k': 31, 'l': 37, 'm': 41, 'n': 43, 'o': 47, 'p': 53, 'q': 59, 'r': 61, 's': 67, 't': 71, 'u': 79, 'v': 83, 'w': 89, 'x': 97, 'y': 101, 'z': 103}
	if len(strs) == 0 {
		return nil
	}
	m2 := make(map[int64][]string)
	for _, v := range strs {
		var tmp int64
		tmp = 1
		for _, v2 := range v {
			tmp *= m[v2]
		}
		if _, ok := m2[tmp]; ok {
			m2[tmp] = append(m2[tmp], v)
		} else {
			m2[tmp] = []string{v}
		}
	}
	result := make([][]string, 0)
	for _, v := range m2 {
		tmp := make([]string, len(v))
		copy(tmp, v)
		result = append(result, tmp)
	}
	return result
}
