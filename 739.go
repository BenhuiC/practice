package partice

func dailyTemperatures(T []int) []int {
	result := make([]int, len(T), len(T))
	if T == nil || len(T) == 0 {
		return result
	}
	st := make(stack, 0)
	for i, v := range T {
		if st.Empty() || T[st.Peek()] > v {
			st = append(st, i)
			continue
		}
		for !st.Empty() && T[st.Peek()] < v {
			pos := st.Pop()
			result[pos] = i - pos
		}
		st = append(st, i)
	}

	return result
}
