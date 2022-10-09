package partice

func scoreOfParentheses(s string) int {
	bal := 0
	res := 0
	for i, v := range s {
		if v == '(' {
			bal++
		} else {
			bal--
			if s[i-1] == '(' {
				res += 1 << bal
			}
		}
	}
	return res
}

// (()(()))
