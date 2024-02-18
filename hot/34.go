package hot

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 || target < nums[0] || target > nums[n-1] {
		return []int{-1, -1}
	}
	left, right := -1, -1

	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			r = mid - 1
			if mid == 0 || nums[mid-1] != target {
				left = mid
				break
			}
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if left == -1 {
		return []int{-1, -1}
	}

	l, r = 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			l = mid + 1
			if mid == n-1 || nums[mid+1] != target {
				right = mid
				break
			}
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	res := []int{left, right}
	return res
}
