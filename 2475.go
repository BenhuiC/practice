package partice

import "sort"

func unequalTriplets(nums []int) int {
	var res2475 int
	sort.Ints(nums)
	n := len(nums)
	for i, j := 0, 0; i < n; i = j {
		for j = i + 1; j < n && nums[j] == nums[i]; j++ {
		}
		res2475 += (j - i) * i * (n - j)
	}

	return res2475
}

/*
* 排序对结果没有影响
1. 先对数组进行排序，这样相同值的数字就会排在一起
2. 设[i,j)为相同值的数字的区间，那么这个区间对应的三元组数量为i * (j-i) * (n-j)
即：区间左边的数字个数 * 区间内相同数字的个数 * 区间右边的数字个数

2 3 4 4 4 5
i=2 j=5
n=6
*/
