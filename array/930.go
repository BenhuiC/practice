package array

func numSubarraysWithSum(nums []int, goal int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	res := 0
	sumMap := map[int]int{0: 1}
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		res += sumMap[sum-goal]
		sumMap[sum]++
	}

	return res
}
