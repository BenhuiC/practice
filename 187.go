package partice

func findRepeatedDnaSequences(s string) []string {
	res187 := make([]string, 0)
	if len(s) <= 10 {
		return res187
	}
	m := make(map[string]bool)
	add := make(map[string]bool)
	for i := 0; i <= len(s)-10; i++ {
		tmp := s[i : i+10]
		if m[tmp] && !add[tmp] {
			res187 = append(res187, tmp)
			add[tmp] = true
		}
		m[tmp] = true
	}
	return res187
}
