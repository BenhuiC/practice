package partice

func longestValidParentheses(s string) int {
	sta := make(stack, 0)
	ary := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			sta = append(sta, i)
		} else {
			if sta.Empty() {
				continue
			}
			t := sta.Pop()
			ary[t] = true
			ary[i] = true
		}
	}
	var result, tmp int
	for _, v := range ary {
		if v {
			tmp++
			if tmp > result {
				result = tmp
			}
		} else {
			tmp = 0
		}
	}
	return result
}
