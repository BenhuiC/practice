package dp

import "sort"

func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	maxV := rewardValues[len(rewardValues)-1]
	dp := make([]int, 2*maxV)
	dp[0] = 1
	for i, v := range rewardValues {
		if i > 0 && rewardValues[i] == rewardValues[i-1] {
			continue
		}
		for k := 2*v - 1; k >= v; k-- {
			if dp[k-v] == 1 {
				dp[k] = 1
			}
		}
	}

	res := 0
	for i := 2*maxV - 1; i >= 0; i-- {
		if dp[i] == 1 {
			res = i
			break
		}
	}
	return res
}
