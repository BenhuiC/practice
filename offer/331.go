package offer

import "strings"

func isValidSerialization(preorder string) bool {
	stack := make([]int, 0)
	ary := strings.Split(preorder, ",")
	if len(ary) <= 0 {
		return false
	}
	stack = append(stack, 1)
	for _, v := range ary {
		if len(stack) == 0 {
			return false
		}

		if v != "#" {
			stack = append(stack, 2)
		} else {
			stack[len(stack)-1]--
			for len(stack) > 0 && stack[len(stack)-1] <= 0 {
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					stack[len(stack)-1]--
				}
			}
		}
	}
	return len(stack) == 0
}
