package partice

func largestRectangleArea(heights []int) int {
	var result int
	if len(heights) <= 0 {
		return 0
	}
	tp := []int{0}
	tp = append(tp, heights...)
	tp = append(tp, 0)
	st := make(stack, 0)
	for i, v := range tp {
		for !st.Empty() && tp[st.Peek()] > v {
			h := tp[st.Pop()]
			area := (i - st.Peek() - 1) * h
			if area > result {
				result = area
			}
		}
		st = append(st, i)
	}

	return result
}
