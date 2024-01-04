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
	dp := make([][2]int, mx+1)
	for i := 1; i <= mx; i++ {
		v, ok := mp[i]
		if !ok {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1])
			dp[i][1] = dp[i][0]
		} else {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1])
			if i == 1 {
				dp[i][1] = v
			} else {
				dp[i][1] = max(dp[i-2][0], dp[i-2][1]) + v
			}
		}
	}

	return max(dp[mx][0], dp[mx][1])
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
