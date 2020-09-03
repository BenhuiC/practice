package partice

import "sort"

var res [][]int

func subsetsWithDup(nums []int) [][]int {
	res = make([][]int, 0)
	if len(nums) <= 0 {
		return res
	}
	a := make([]int, 0)
	sort.Ints(nums)
	ress(nums, a, 0)
	return res
}

func ress(nums, a []int, n int) {
	tmp := make([]int, len(a))
	copy(tmp, a)
	res = append(res, tmp)
	for i := n; i < len(nums); i++ {
		if i-1 >= n && nums[i-1] == nums[i] {
			continue
		}
		a = append(a, nums[i])
		ress(nums, a, i+1)
		a = a[:len(a)-1]
	}
}
