package offer

import (
	"fmt"
	"math"
	"strconv"
)

func nextGreaterElement(n int) int {
	ary := []byte(fmt.Sprint(n))
	i := len(ary) - 2
	for i >= 0 && ary[i] >= ary[i+1] {
		i--
	}
	if i < 0 {
		return -1
	}
	j := len(ary) - 1
	for j >= i && ary[j] <= ary[i] {
		j--
	}
	ary[i], ary[j] = ary[j], ary[i]

	l, r := i+1, len(ary)-1
	for l < r {
		ary[l], ary[r] = ary[r], ary[l]
		l++
		r--
	}

	res, _ := strconv.Atoi(string(ary))
	if res > math.MaxInt32 {
		return -1
	}
	return res
}
