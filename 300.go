package partice

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	ary := make([]int, len(nums))
	max := 0
	for _, v := range nums {
		l, h := 0, max
		for l < h {
			mid := (l + h) / 2
			if ary[mid] < v {
				l = mid + 1
			} else {
				h = mid
			}
		}
		ary[l] = v
		if l == max {
			max++
		}
	}
	return max
}
