package partice

func minSteps(s string, t string) int {
	if s == t {
		return 0
	}
	res := 0
	m := map[uint8]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		m[t[i]]--
	}

	for _, v := range m {
		if v > 0 {
			res += v
		}
	}
	return res
}
