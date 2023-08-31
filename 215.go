package partice

func findKthLargest(nums []int, k int) int {
	l, r := 0, len(nums)-1
	val := nums[0]
	for l < r {
		for l < r && nums[r] <= val {
			r--
		}
		nums[l] = nums[r]
		for l < r && nums[l] > val {
			l++
		}
		nums[r] = nums[l]
	}
	nums[l] = val
	if k == l+1 {
		return val
	} else if k < l+1 {
		return findKthLargest(nums[:l], k)
	} else {
		return findKthLargest(nums[l+1:], k-l-1)
	}
}
