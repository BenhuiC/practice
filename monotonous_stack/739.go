package monotonous_stack

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	ary := make([]int, 0)
	for i := 0; i < len(temperatures); i++ {
		if len(ary) == 0 {
			ary = append(ary, i)
			continue
		}
		for len(ary) != 0 && temperatures[ary[len(ary)-1]] < temperatures[i] {
			idx := ary[len(ary)-1]
			res[idx] = i - idx
			ary = ary[:len(ary)-1]
		}
		ary = append(ary, i)
	}
	return res
}
