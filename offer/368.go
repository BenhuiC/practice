package offer

import "sort"

func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	sort.Ints(nums)
	dp := make([]int, n)
	par := make([]int, n)
	maxIdx, maxVal := n-1, 1
	for i := n - 1; i >= 0; i-- {
		idx := 0
		for j := i + 1; j < n; j++ {
			if nums[j]%nums[i] == 0 && dp[j] > dp[idx] {
				idx = j
				dp[i] = dp[j] + 1
				par[i] = j
				if dp[i] > maxVal {
					maxIdx = i
					maxVal = dp[i]
				}
			}
		}
		if idx == 0 {
			par[i] = -1
			dp[i] = 1
		}
	}
	res := make([]int, 0, maxVal)
	for p := maxIdx; p != -1; p = par[p] {
		res = append(res, nums[p])
	}
	return res
}
