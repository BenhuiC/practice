package partice

func validateStackSequences(pushed []int, popped []int) bool {
	res := false
	stack := make([]int, 0, len(pushed)/2)
	i, j := 0, 0
	n := len(pushed)
	for j < n {
		if len(stack) == 0 || stack[len(stack)-1] != popped[j] {
			if i >= n {
				break
			}
			stack = append(stack, pushed[i])
			i++
		}
		if stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}

	res = len(stack) == 0

	return res
}

/*
1 2 3 4 5
5 4 3 2 1
*/
