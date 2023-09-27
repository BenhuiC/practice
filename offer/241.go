package offer

import (
	"strconv"
)

func diffWaysToCompute(expression string) []int {
	nums := make([]int, 0)
	op := make([]byte, 0)
	i, j := 0, 0
	for i < len(expression) && j < len(expression) {
		for j = i + 1; j < len(expression) && expression[j] >= '0' && expression[j] <= '9'; j++ {
		}
		nu, _ := strconv.Atoi(expression[i:j])
		nums = append(nums, nu)
		if j < len(expression) {
			op = append(op, expression[j])
		}
		i = j + 1
	}

	var split func(l, r int) []int
	split = func(l, r int) []int {
		if l > r || l >= len(nums) {
			return nil
		}
		if l == r {
			return []int{nums[l]}
		}
		res := make([]int, 0)
		for i := l; i < r; i++ {
			left := split(l, i)
			right := split(i+1, r)
			operate := op[i]
			for _, lval := range left {
				for _, rval := range right {
					switch operate {
					case '-':
						res = append(res, lval-rval)
					case '+':
						res = append(res, lval+rval)
					case '*':
						res = append(res, lval*rval)
					}
				}
			}
		}
		return res
	}

	return split(0, len(nums)-1)
}
