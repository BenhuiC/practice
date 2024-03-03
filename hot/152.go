package hot

func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	res := nums[0]
	minV, maxV := nums[0], nums[0] // 对应i之前最小值和最大值
	for i := 1; i < n; i++ {
		tmpMin, tmpMax := minV, maxV
		minV = min(nums[i], min(tmpMin*nums[i], tmpMax*nums[i]))
		maxV = max(nums[i], max(tmpMin*nums[i], tmpMax*nums[i]))
		res = max(res, maxV)
	}

	return res
}
