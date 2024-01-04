package dp

import "math"

func deleteAndEarn(nums []int) int {
	mp := make(map[int]int)
	mx := 0
	for _, n := range nums {
		if n > mx {
			mx = n
		}
		mp[n] += n
	}
	dp := make([]int, mx+1)
	for i := 1; i <= mx; i++ {
		v, ok := mp[i]
		if !ok {
			dp[i] = dp[i-1]
		} else {
			if i == 1 {
				dp[i] = v
			} else {
				dp[i] = max(dp[i-2]+v, dp[i-1])
			}
		}
	}

	return dp[mx]
}

func max(a ...int) int {
	mx := math.MinInt
	for _, v := range a {
		if v > mx {
			mx = v
		}
	}
	return mx
}
