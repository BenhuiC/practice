package partice

import "math"

func maxSumDivThree(nums []int) int {
	// 使用动态规划的思路，dp数组保存对于当前处理的数字余j时的最大值
	dp := []int{0, math.MinInt32, math.MinInt32}
	tmp := make([]int, 3)
	for _, v := range nums {
		for j := 0; j < 3; j++ {
			tmp[j] = Max(dp[j], dp[(j+3-v%3)%3]+v)
		}
		copy(dp, tmp)
	}

	return dp[0]
}
