package partice

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	all := 0
	for _, v := range nums {
		all += v
	}
	if all%k != 0 {
		return false
	}
	avg := all / k
	sort.Ints(nums)
	if nums[len(nums)-1] > avg {
		return false
	}
	n := len(nums)
	mask := 1<<n - 1        // 最终结果，每一位都是1
	st := make([]int, 1<<n) // 状态数组 -1 不符合，1 符合，0 初始化状态

	var dfs func(int, int) bool
	dfs = func(s int, v int) bool {
		// s是当前数组的使用状态，通过二进制位来判断。第i位为1则nums[i]已使用
		if s == mask {
			return true
		}
		if st[s] != 0 {
			return st[s] == 1
		}
		for i, val := range nums {
			if s>>i&1 == 1 { // 当前位已使用
				continue
			}
			if v+val > avg {
				break
			}
			if dfs(s|1<<i, (v+val)%avg) {
				st[s] = 1
				return true
			}
		}
		st[s] = -1
		return false
	}

	return dfs(0, 0)
}
