package offer

func minLength(s string) int {
	stack := make([]byte, 0)
	for i := range s {
		stack = append(stack, s[i])
		if l := len(stack); l >= 2 && (string(stack[l-2:l]) == "AB" || string(stack[l-2:l]) == "CD") {
			stack = stack[:l-2]
		}
	}
	return len(stack)
}
