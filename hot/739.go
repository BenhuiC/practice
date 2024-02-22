package hot

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	if n < 2 {
		return res
	}

	stk := make([]int, 0)
	for i, t := range temperatures {
		if len(stk) == 0 || temperatures[stk[len(stk)-1]] >= t {
			stk = append(stk, i)
			continue
		}
		for len(stk) > 0 && temperatures[stk[len(stk)-1]] < t {
			last := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			res[last] = i - last
		}
		stk = append(stk, i)
	}

	return res
}
