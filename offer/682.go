package offer

import "strconv"

func calPoints(operations []string) int {
	stack := make([]int, 0)
	res := 0
	for _, op := range operations {
		switch op {
		case "C":
			t := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res -= t
		case "D":
			t := stack[len(stack)-1] * 2
			stack = append(stack, t)
			res += t
		case "+":
			n := len(stack)
			t := stack[n-1] + stack[n-2]
			stack = append(stack, t)
			res += t
		default:
			sc, _ := strconv.Atoi(op)
			res += sc
			stack = append(stack, sc)
		}
	}

	return res
}
