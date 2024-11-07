package offer150

func rotate(nums []int, k int) {
	n := len(nums)
	if k >= n {
		k %= n
	}
	reverse := func(left, right int) {
		for left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	reverse(0, n-1)
	reverse(0, k-1)
	reverse(k, n-1)
}
