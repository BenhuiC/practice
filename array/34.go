package array

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	l, h := 0, len(nums)-1
	var mid int
	for l <= h {
		mid = l + (h-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			h = mid - 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				res[0] = mid
				break
			}
			h = mid - 1
		}
	}
	if res[0] == -1 {
		return res
	}

	l, h = mid, len(nums)-1
	for l <= h {
		mid = l + (h-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			h = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != target {
				res[1] = mid
				break
			}
			l = mid + 1
		}
	}

	return res
}
