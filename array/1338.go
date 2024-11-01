package array

import "sort"

func minSetSize(arr []int) int {
	m := make(map[int]int)
	n := len(arr)
	for _, v := range arr {
		m[v] += 1
	}
	nums := make([]int, 0, len(m))
	for _, v := range m {
		nums = append(nums, v)
	}
	sort.Ints(nums)

	half := n / 2
	if n&1 == 1 {
		half += 1
	}
	res := 0
	for i := len(nums) - 1; i >= 0; i-- {
		res += 1
		half -= nums[i]
		if half <= 0 {
			break
		}
	}
	return res
}
