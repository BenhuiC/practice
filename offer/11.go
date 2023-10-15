package offer

func maxArea(height []int) int {
	res := 0
	left, right := 0, len(height)-1
	for left < right {
		var area int
		if height[left] < height[right] {
			area = height[left] * (right - left)
			left++
		} else {
			area = height[right] * (right - left)
			right--
		}
		if area > res {
			res = area
		}
	}

	return res
}
