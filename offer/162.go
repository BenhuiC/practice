package offer

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)>>1
		if mid < len(nums)-1 && nums[mid+1] >= nums[mid] {
			l++
		} else if mid > 0 && nums[mid-1] >= nums[mid] {
			r--
		} else {
			return mid
		}
	}
	return l
}
