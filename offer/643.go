package offer

func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	res := nums[0]
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
		if i == k-1 {
			res = nums[i]
		} else if i > k-1 && nums[i]-nums[i-k] > res {
			res = nums[i] - nums[i-k]
		}
	}
	return float64(res) / float64(k)
}
