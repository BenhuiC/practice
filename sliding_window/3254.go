package sliding_window

func resultsArray(nums []int, k int) []int {
	n := len(nums)
	res := make([]int, n-k+1)
	left, right := 0, 0
	last := 0 // 上次开始递增的下标
	for right < n {
		if right < left+k-1 && right < n {
			right++
			if nums[right]-nums[right-1] != 1 {
				last = right
			}
			continue
		}

		if last <= left {
			res[left] = nums[right]
		} else {
			res[left] = -1
		}
		left++
		right++
		if right < n && nums[right]-nums[right-1] != 1 {
			last = right
		}
	}
	return res
}
