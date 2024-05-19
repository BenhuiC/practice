package array

func removeDuplicates1047(s string) string {
	stk := make([]byte, 0)
	for i := range s {
		if len(stk) == 0 || s[i] != stk[len(stk)-1] {
			stk = append(stk, s[i])
		} else {
			stk = stk[:len(stk)-1]
		}
	}
	return string(stk)
}
