package hot

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)

	leftMax, rightMax := make([]int, n), make([]int, n)

	for i, v := range nums {
		if i%k == 0 {
			leftMax[i] = v
		} else {
			leftMax[i] = max(leftMax[i-1], v)
		}
	}

	for j := n - 1; j >= 0; j-- {
		if j == n-1 || (j+1)%k == 0 {
			rightMax[j] = nums[j]
		} else {
			rightMax[j] = max(rightMax[j+1], nums[j])
		}
	}

	res := make([]int, n-k+1)
	for i := 0; i <= n-k; i++ {
		res[i] = max(rightMax[i], leftMax[i+k-1])
	}

	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
