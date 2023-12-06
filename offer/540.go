package offer

func singleNonDuplicate(nums []int) int {
	n := len(nums)
	if n&1 == 0 {
		return -1
	}
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if mid == 0 || mid == n-1 {
			return nums[mid]
		}
		if nums[mid] != nums[mid-1] && nums[mid] != nums[mid+1] {
			return nums[mid]
		}
		if mid&1 == 0 {
			if nums[mid] == nums[mid+1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if nums[mid] == nums[mid+1] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}
