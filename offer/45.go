package offer

import (
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	res := ""
	if len(nums) == 0 {
		return ""
	}
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] == nums[j] {
			return true
		}
		strI := strconv.Itoa(nums[i])
		strJ := strconv.Itoa(nums[j])
		if strI+strJ < strJ+strI {
			return true
		}
		return false
	})

	for _, v := range nums {
		res += strconv.Itoa(v)
	}

	return res
}

/*
121 12
212 21

13241321324
13241324132

*/
