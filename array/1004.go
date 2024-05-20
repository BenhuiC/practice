package array

func longestOnes(nums []int, k int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	res := 0
	leftSum, rightSum := 0, 0
	left := 0
	for right, v := range nums {
		if v == 0 {
			rightSum++
		}
		for rightSum-leftSum > k {
			if nums[left] == 0 {
				leftSum++
			}
			left++
		}
		res = max(res, right-left+1)
	}

	return res
}
