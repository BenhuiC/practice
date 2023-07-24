package partice

/*
答案存在两种情况
1. nums[i:j] i>=0,j<n
2. nums[0:i],nums[j:n]
*/
func maxSubarraySumCircular(nums []int) int {
	res918, prev, leftSum := nums[0], nums[0], nums[0]
	n := len(nums)
	// leftMax为从0开始到当前索引的最大和
	leftMax := make([]int, n)
	leftMax[0] = nums[0]
	for i := 1; i < n; i++ {
		leftSum += nums[i]
		prev = Max(prev+nums[i], nums[i])
		res918 = Max(res918, prev)
		leftMax[i] = Max(leftMax[i-1], leftSum)
	}

	rightSum := 0
	// 再从右到左计算出最大值
	for j := n - 1; j > 0; j-- {
		rightSum += nums[j]
		res918 = Max(res918, leftMax[j-1]+rightSum)
	}

	return res918
}
