package hot

func decodeString(s string) string {
	if len(s) <= 1 {
		return s
	}
	var curNum int
	var curStr string
	numStk := make(stack[int], 0)
	strStk := make(stack[string], 0)
	for _, c := range s {
		if c == '[' {
			numStk = append(numStk, curNum)
			strStk = append(strStk, curStr)
			curNum = 0
			curStr = ""
		} else if c >= '0' && c <= '9' {
			curNum = curNum*10 + int(c-'0')
		} else if c >= 'a' && c <= 'z' {
			curStr += string(c)
		} else {
			n := numStk.pop()
			str := strStk.pop()
			for i := 0; i < n; i++ {
				str += curStr
			}
			curStr = str
		}
	}

	return curStr
}

type stack[T any] []T

func (s *stack[T]) push(val T) {
	*s = append(*s, val)
}

func (s *stack[T]) pop() (val T) {
	if len(*s) == 0 {
		return
	}
	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r
}
