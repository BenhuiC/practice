package dp

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	res := 1
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
			res = Max(res, dp[i])
		}
	}

	return res
}
