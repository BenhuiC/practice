package hot

func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n < 1 {
		return 0
	}
	res := 0
	heights = append([]int{0}, append(heights, 0)...)
	stk := make([]int, 0)
	for i, h := range heights {
		for len(stk) > 0 && heights[stk[len(stk)-1]] > h {
			t := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			area := (i - stk[len(stk)-1] - 1) * heights[t]
			if area > res {
				res = area
			}
		}

		stk = append(stk, i)
	}

	return res
}
