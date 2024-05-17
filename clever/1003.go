package clever

func isValid(s string) bool {
	n := len(s)
	if n < 3 {
		return false
	}
	if s[0] != 'a' || s[n-1] != 'c' {
		return false
	}
	stack := make([][]int32, 0)
	for _, c := range s {
		if c == 'a' {
			stack = append(stack, []int32{'a'})
			continue
		}
		if len(stack) == 0 {
			return false
		} else {
			top := stack[len(stack)-1]
			if c-top[len(top)-1] == 1 {
				top = append(top, c)
				if len(top) == 3 {
					stack = stack[:len(stack)-1]
				} else {
					stack[len(stack)-1] = top
				}

				continue
			}
			return false
		}
	}
	return len(stack) == 0
}

func isValid2(s string) bool {
	stack := make([]int32, 0)
	for _, c := range s {
		stack = append(stack, c)
		n := len(stack)
		if n >= 3 && string(stack[n-3:]) == "abc" {
			stack = stack[:n-3]
		}
	}
	return len(stack) == 0
}
