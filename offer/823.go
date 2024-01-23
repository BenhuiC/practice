package offer

import "sort"

func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)
	n := len(arr)
	const mod int = 1e9 + 7
	res := 0
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		for l, r := 0, i-1; l <= r; l++ {
			for l <= r && arr[l]*arr[r] > arr[i] {
				r--
			}
			if l <= r && arr[l]*arr[r] == arr[i] {
				if l == r {
					dp[i] = (dp[i] + dp[l]*dp[r]) % mod
				} else {
					dp[i] = (dp[i] + 2*dp[l]*dp[r]) % mod
				}
			}
		}
		res = (res + dp[i]) % mod
	}

	return res
}
