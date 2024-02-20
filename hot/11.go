package hot

func maxArea(height []int) int {
	n := len(height)
	if n < 2 {
		return 0
	}
	res := 0
	l, r := 0, n-1
	/*
		此时r-l作为底，底最大，在向内收缩时，判断两条边，矮的一边向内收缩
	*/
	for l < r {
		var tmp int
		if height[l] < height[r] {
			tmp = (r - l) * height[l]
			l++
		} else {
			tmp = (r - l) * height[r]
			r--
		}
		if tmp > res {
			res = tmp
		}
	}

	return res
}
