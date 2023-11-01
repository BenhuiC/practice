package offer

import "strings"

func lengthLongestPath(input string) int {
	ary := strings.Split(input, "\n")
	res := 0
	var baseLen int
	stack := make([]int, 0)
	for _, v := range ary {
		items := strings.Split(v, "\t")
		if len(items) == 0 {
			continue
		}
		curLevel := len(items) - 1
		for len(stack) > curLevel {
			baseLen -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		item := items[len(items)-1]
		// file
		if strings.Contains(item, ".") {
			if baseLen+len(item) > res {
				res = baseLen + len(item)
			}
			continue
		}
		//dir
		baseLen += len(item) + 1
		stack = append(stack, len(item)+1)
	}
	return res
}
