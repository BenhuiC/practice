package offer

import "sort"

func numFriendRequests(ages []int) int {
	sort.Ints(ages)

	res := 0
	l, r := 0, 0
	for _, v := range ages {
		if v < 15 {
			continue
		}
		for ages[l]*2 <= v+14 {
			l++
		}
		for r+1 < len(ages) && ages[r+1] <= v {
			r++
		}
		res += r - l
	}

	return res
}
