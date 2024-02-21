package hot

func isValid(s string) bool {
	st := make([]byte, 0)
	for _, v := range s {
		switch v {
		case '(', '[', '{':
			st = append(st, byte(v))
		case ')':
			if len(st) == 0 || st[len(st)-1] != '(' {
				return false
			}
			st = st[:len(st)-1]
		case ']':
			if len(st) == 0 || st[len(st)-1] != '[' {
				return false
			}
			st = st[:len(st)-1]
		case '}':
			if len(st) == 0 || st[len(st)-1] != '{' {
				return false
			}
			st = st[:len(st)-1]
		}
	}
	return len(st) == 0
}
