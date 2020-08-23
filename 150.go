package partice

import "strconv"

func evalRPN(tokens []string) int {
	if tokens == nil || len(tokens) == 0 {
		return 0
	}
	s := make(stack, 0)
	for _, v := range tokens {
		switch v {
		case "+":
			a := s.Pop() + s.Pop()
			s = append(s, a)
		case "-":
			a := -s.Pop() + s.Pop()
			s = append(s, a)
		case "*":
			a := s.Pop() * s.Pop()
			s = append(s, a)
		case "/":
			a, b := s.Pop(), s.Pop()
			s = append(s, b/a)
		default:
			val, _ := strconv.ParseInt(v, 10, 64)
			s = append(s, int(val))
		}
	}
	return s.Pop()
}
