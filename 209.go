package partice

func minSubArrayLen(target int, nums []int) int {
	res209 := len(nums) + 1
	left, right := 0, 0
	sum := 0
	for right < len(nums) {
		sum += nums[right]
		right++
		if sum < target {
			continue
		}
		for sum >= target && left < right {
			if right-left < res209 {
				res209 = right - left
			}
			sum -= nums[left]
			left++
		}
	}
	if res209 == len(nums)+1 {
		return 0
	}
	return res209
}
