package hot

func searchInsert(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if target <= nums[0] {
		return 0
	}
	if target > nums[n-1] {
		return n
	}
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if target == nums[mid] {
			return mid
		} else if nums[mid] > target {
			if nums[mid-1] < target {
				return mid
			}
			r = mid - 1
		} else {
			if nums[mid+1] >= target {
				return mid + 1
			}
			l = mid + 1
		}
	}
	return l
}
