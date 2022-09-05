package partice

import "sort"

var res90 [][]int

func subsetsWithDup(nums []int) [][]int {
	res90 = make([][]int, 0)
	if len(nums) <= 0 {
		return res90
	}
	a := make([]int, 0)
	sort.Ints(nums)
	ress(nums, a, 0)
	return res90
}

func ress(nums, a []int, n int) {
	tmp := make([]int, len(a))
	copy(tmp, a)
	res90 = append(res90, tmp)
	for i := n; i < len(nums); i++ {
		if i-1 >= n && nums[i-1] == nums[i] {
			continue
		}
		a = append(a, nums[i])
		ress(nums, a, i+1)
		a = a[:len(a)-1]
	}
}
