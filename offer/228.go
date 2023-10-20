package offer

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	res := make([]string, 0)
	mn := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			continue
		}
		if mn == nums[i-1] {
			res = append(res, fmt.Sprint(mn))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", mn, nums[i-1]))
		}
		mn = nums[i]
	}

	if mn == nums[len(nums)-1] {
		res = append(res, fmt.Sprint(mn))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", mn, nums[len(nums)-1]))
	}

	return res
}
