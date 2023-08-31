package partice

func findPeakElement(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	if len(nums) == 2 {
		if nums[1] > nums[0] {
			return 1
		}
		return 0
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if mid == 0 && nums[mid] > nums[mid+1] {
			return 0
		}
		if mid == len(nums)-1 && nums[mid] > nums[mid-1] {
			return mid
		}
		if nums[mid-1] < nums[mid] && nums[mid] > nums[mid+1] {
			return mid
		}
		if l != mid && nums[l] >= nums[mid] {
			r = mid - 1
			continue
		}
		l = mid + 1
	}

	return 0
}
