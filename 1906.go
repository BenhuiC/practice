package partice

import "math"

/*
2 <= nums.length <= 105
1 <= nums[i] <= 100
1 <= queries.length <= 2 * 104
0 <= li < ri < nums.length
*/
func minDifference(nums []int, queries [][]int) []int {
	res1906 := make([]int, 0, len(queries))
	n := len(nums)
	// prefix 数组保存前缀和，prefix[i]保存的是nums[:i]中数字的出现次数
	prefix := make([][]int, n+1)
	prefix[0] = make([]int, 101)
	for i := 0; i < n; i++ {
		prefix[i+1] = make([]int, 101)
		copy(prefix[i+1], prefix[i])
		prefix[i+1][nums[i]]++
	}

	for i := range queries {
		left, right := queries[i][0], queries[i][1]
		last, best := 0, math.MaxInt
		for j := 1; j <= 100; j++ {
			// 这样通过比较prefix[right+1]和prefix[left]就可以得知数字是否在nums[left:right+1]中出现
			if prefix[left][j] != prefix[right+1][j] {
				// 如果差值等于查询区间的长度，则表示区间内是同样的数字，则返回-1
				if prefix[right+1][j]-prefix[left][j] == right-left+1 {
					best = -1
					break
				}
				if last != 0 {
					// 通过计算j-last和best的最小值，来更新best
					// last是从1到100，所以不需要再进行排序，last->j之间的数字肯定没有出现过
					best = Min(best, j-last)
				}
				last = j
			}
		}
		res1906 = append(res1906, best)
	}

	return res1906
}
