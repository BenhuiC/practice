package offer

func trap(height []int) int {
	nl := len(height)
	leftMax, rightMax := make([]int, nl), make([]int, nl)
	for i := 0; i < nl; i++ {
		leftMax[i] = height[i]
		if i > 0 && leftMax[i-1] > height[i] {
			leftMax[i] = leftMax[i-1]
		}

		rightMax[nl-1-i] = height[nl-1-i]
		if i > 0 && rightMax[nl-i] > height[nl-1-i] {
			rightMax[nl-1-i] = rightMax[nl-i]
		}
	}

	res := 0
	for i := 1; i < nl-1; i++ {
		t := min(leftMax[i], rightMax[i]) - height[i]
		if t > 0 {
			res += t
		}
	}

	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
