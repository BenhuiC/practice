package offer

func nextPermutation(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	reverse := func(ary []int) {
		l, r := 0, len(ary)-1
		for l < r {
			ary[l], ary[r] = ary[r], ary[l]
			l++
			r--
		}
	}
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i < 0 {
		reverse(nums)
		return
	}
	j := n - 1
	for j > i && nums[j] <= nums[i] {
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]
	reverse(nums[i+1:])
}
