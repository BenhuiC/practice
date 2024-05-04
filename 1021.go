package practice

func removeOuterParentheses(s string) string {
	res := ""
	idx := 0
	left := 0
	for idx < len(s) {
		i := idx
		for ; i < len(s); i++ {
			if s[i] == '(' {
				left++
			} else {
				left--
			}
			if left == 0 {
				res += s[idx+1 : i]
				break
			}
		}
		idx = i
	}
	return res
}
