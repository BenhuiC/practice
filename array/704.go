package array

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if target < nums[0] || target > nums[n-1] {
		return -1
	}
	l, r := 0, n-1
	var mid int
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
