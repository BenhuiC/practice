package offer

func countSegments(s string) int {
	res := 0
	for i := 0; i < len(s); {
		if s[i] == ' ' {
			i++
			continue
		}
		j := i
		for j < len(s) && s[j] != ' ' {
			j++
		}
		i = j
		res++
	}
	return res
}
