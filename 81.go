package partice

func search2(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid] < nums[high] {
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else if nums[mid] > nums[high] {
			if target < nums[mid] && target >= nums[low] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			high--
		}
	}
	return false
}
