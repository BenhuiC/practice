package hot

import "math"

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 1 {
		return false
	}
	sum := 0
	maxV := math.MinInt
	for _, v := range nums {
		sum += v
		if v > maxV {
			maxV = v
		}
	}
	if sum&1 == 1 {
		return false
	}
	target := sum / 2
	if maxV > target {
		return false
	}
	if target == maxV {
		return true
	}

	dp := make([]bool, target+1)
	dp[0] = true
	for _, v := range nums {
		for j := target; j >= v; j-- {
			dp[j] = dp[j] || dp[j-v]
		}
	}
	return dp[target]
}
