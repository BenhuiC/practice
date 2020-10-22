package partice

func searchRange(nums []int, target int) []int {
	l, h := -1, -1
	low, high := 0, len(nums)
	for low < high {
		mid := (high-low)/2 + low
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid
		} else {
			if mid+1 >= len(nums) || nums[mid+1] > target {
				h = mid
				high = mid
			}
			if mid-1 < 0 || nums[mid-1] < target {
				l = mid
				break
			} else {
				high = mid
			}
		}
	}
	if l == -1 {
		return []int{-1, -1}
	}
	if h != -1 {
		return []int{l, h}
	}

	low, high = 0, len(nums)
	for low < high {
		mid := (high-low)/2 + low
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid
		} else {
			if mid+1 >= len(nums) || nums[mid+1] > target {
				h = mid
				break
			} else {
				low = mid + 1
			}
		}
	}

	return []int{l, h}
}
