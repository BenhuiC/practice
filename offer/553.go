package offer

import (
	"fmt"
	"strconv"
	"strings"
)

func optimalDivision(nums []int) string {
	n := len(nums)
	if n == 1 {
		return strconv.Itoa(nums[0])
	}
	if n == 2 {
		return fmt.Sprintf("%d/%d", nums[0], nums[1])
	}
	ans := &strings.Builder{}
	ans.WriteString(fmt.Sprintf("%d/(%d", nums[0], nums[1]))
	for _, num := range nums[2:] {
		ans.WriteByte('/')
		ans.WriteString(strconv.Itoa(num))
	}
	ans.WriteByte(')')
	return ans.String()
}

func optimalDivision2(nums []int) string {
	var max func(ary []int) (float64, string)
	var min func(ary []int) (float64, string)
	max = func(ary []int) (float64, string) {
		if len(ary) == 0 {
			return 0, ""
		}
		m := float64(ary[0])
		s := fmt.Sprintf("%d", ary[0])
		for i := 1; i < len(ary); i++ {
			m /= float64(ary[i])
			s = fmt.Sprintf("%s/%d", s, ary[i])
		}
		for i := 1; i < len(ary)-1; i++ {
			lmx, ls := max(ary[:i])
			rmi, rl := min(ary[i:])
			t := lmx / rmi
			if t > m {
				m = t
				s = fmt.Sprintf("%s/(%s)", ls, rl)
			}
		}
		return m, s
	}
	min = func(ary []int) (float64, string) {
		if len(ary) == 0 {
			return 0, ""
		}
		m := float64(ary[0])
		s := fmt.Sprint(ary[0])
		for i := 1; i < len(ary); i++ {
			m /= float64(ary[i])
			s = fmt.Sprintf("%s/%d", s, ary[i])
		}

		for i := 1; i < len(ary)-1; i++ {
			lmi, ls := min(ary[:i])
			rmx, rs := max(ary[i:])
			t := lmi / rmx
			if t < m {
				m = t
				s = fmt.Sprintf("%s/(%s)", ls, rs)
			}
		}
		return m, s
	}
	_, res := max(nums)
	return res
}
