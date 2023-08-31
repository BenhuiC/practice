package interview

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	n := len(s1)
	m := make(map[uint8]int)
	for i := 0; i < n; i++ {
		m[s1[i]]++
		m[s2[i]]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
