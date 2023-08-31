package partice

import "sort"

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	res1498 := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		l, r := i, n
		mid := (l + r) / 2
		for l < r {
			if nums[i]+nums[mid] > target {
				r = mid - 1
				continue
			}
			if mid == n-1 || nums[i]+nums[mid+1] > target {
				res1498 += mid - i + 1
				break
			} else {
				l = mid
			}
		}
	}
	return res1498
}
