package partice

func isValid(s string) bool {
	sta := make([]int32, 0)
	m := map[int32]int32{')': '(', '}': '{', ']': '['}
	for _, v := range s {
		switch v {
		case '(', '[', '{':
			sta = append(sta, v)
		case ')', ']', '}':
			if len(sta) == 0 || sta[len(sta)-1] != m[v] {
				return false
			}
			sta = sta[:len(sta)-1]
		}
	}
	return len(sta) == 0
}
