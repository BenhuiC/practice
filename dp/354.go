package dp

import (
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] >= envelopes[j][1]
	})
	n := len(envelopes)
	res := 1
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if envelopes[i][1] > envelopes[j][1] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
		res = Max(dp[i], res)
	}

	return res
}
