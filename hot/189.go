package hot

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	if k == 0 {
		return
	}
	tmp := make([]int, k)
	copy(tmp, nums[n-k:])
	copy(nums[k:], nums[:n-k])
	copy(nums[:k], tmp)
}

func rotate2(nums []int, k int) {
	n := len(nums)
	k = k % n
	if k == 0 {
		return
	}

	reverse := func(l, r int) {
		for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	reverse(0, n-1)
	reverse(0, k-1)
	reverse(k, n-1)
}
