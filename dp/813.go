package dp

func largestSumOfAverages(nums []int, k int) float64 {
	n := len(nums)
	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}
	dp := make([][]float64, k)
	for i := 0; i < k; i++ {
		dp[i] = make([]float64, n)
	}
	for i := 0; i < n; i++ {
		dp[0][i] = float64(prefixSum[i+1]) / float64(i+1)
	}

	for i := 1; i < k; i++ {
		for j := 0; j < n; j++ {
			mx := dp[i-1][j]
			for z := 0; z < j; z++ {
				lastLen := j - z
				lastSum := prefixSum[j+1] - prefixSum[z+1]
				mx = maxf(mx, dp[i-1][z]+float64(lastSum)/float64(lastLen))
			}
			dp[i][j] = mx
		}
	}

	return dp[k-1][n-1]
}

func maxf(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
