package partice

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	l, h := 0, len(nums)-1
	for l+1 < h {
		mid := (l + h) / 2
		if nums[mid] > nums[h] {
			l = mid
		} else {
			h = mid
		}
	}
	if nums[l] > nums[h] {
		return nums[h]
	}
	return nums[l]
}
