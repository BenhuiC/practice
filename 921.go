package partice

func minAddToMakeValid(s string) int {
	res := 0
	stk := 0
	for _, v := range s {
		switch v {
		case '(':
			stk++
		case ')':
			if stk <= 0 {
				res++
			} else {
				stk--
			}
		}
	}
	res += stk
	return res
}
