package hot

func findMin(nums []int) int {
	n := len(nums)
	res := 0
	l, r := 0, n-1

	for l <= r {
		mid := l + (r-l)/2
		if (mid == 0 || mid-1 >= 0 && nums[mid] < nums[mid-1]) && (mid == n-1 || mid+1 < n && nums[mid] < nums[mid+1]) {
			res = nums[mid]
			return res
		}
		if nums[mid] < nums[l] {
			r = mid - 1
		} else if nums[mid] > nums[r] {
			l = mid + 1
		} else if nums[mid] > nums[l] && nums[mid] < nums[r] {
			res = nums[l]
			return res
		}
	}

	return res
}
