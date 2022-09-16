package partice

import "sort"

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	dp := make([]int, len(pairs))
	dp[0] = 1
	for i := 1; i < len(pairs); i++ {
		if pairs[i][0] > pairs[i-1][1] {
			dp[i] = dp[i-1] + 1
			continue
		}
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if pairs[i][0] > pairs[j][1] {
				dp[i] = dp[j] + 1
				break
			}
		}
	}
	return dp[len(pairs)-1]
}
