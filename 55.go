package partice

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	mp := make([]bool, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= len(nums)-1 {
			mp[i] = true
			continue
		}
		for _, v := range mp[i+1 : i+nums[i]+1] {
			if v {
				mp[i] = true
				break
			}
		}
	}
	return mp[0]
}

func canJump2(nums []int) bool {
	n := 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] >= n {
			n = 1
		} else {
			n++
		}
	}
	if n > 1 {
		return false
	}
	return true
}
