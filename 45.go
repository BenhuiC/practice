package partice

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	mp := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= len(nums)-1 {
			mp[i] = 1
			continue
		}
		m := mp[i+1]
		for _, v := range mp[i+1 : i+nums[i]+1] {
			if v < m {
				m = v
			}
		}
		mp[i] = m + 1
	}
	return mp[0]
}

func jump2(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	var start, end, max, step int
	for end < len(nums)-1 {
		for i := start; i <= end; i++ {
			if i+nums[i] > max {
				max = i + nums[i]
			}
		}
		start = end + 1
		end = max
		step++
	}
	return step
}
