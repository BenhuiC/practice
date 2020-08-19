package partice

func decodeString(s string) string {
	if s == "" || len(s) == 1 {
		return s
	}
	var num int
	var cur string
	numStack := make(stack, 0)
	strStack := make(stackStr, 0)
	for _, v := range s {
		if v == '[' {
			numStack = append(numStack, num)
			strStack = append(strStack, cur)
			num = 0
			cur = ""
		} else if v >= '0' && v <= '9' {
			num = num*10 + int(v-'0')
		} else if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' {
			cur = cur + string(v)
		} else {
			n := numStack.Pop()
			st := strStack.Pop()
			for i := 0; i < n; i++ {
				st += cur
			}
			cur = st
		}
	}
	return cur
}

type stackStr []string

func (s *stackStr) Pop() string {
	if len(*s) > 0 {
		r := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return r
	}
	return ""
}
