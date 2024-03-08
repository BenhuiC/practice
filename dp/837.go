package dp

func new21Game(n int, k int, maxPts int) float64 {
	// timeout !!
	dp := make([]float64, k+maxPts)
	onePart := 1 / float64(maxPts)
	dp[0] = 1
	for i := 1; i < k+maxPts; i++ {
		for j := 1; j <= maxPts; j++ {
			if i-j >= 0 {
				if i-j < k {
					dp[i] += dp[i-j] * onePart
				}
			} else {
				break
			}
		}
	}

	var kp, np float64
	for i := k; i < k+maxPts; i++ {
		if i <= n {
			np += dp[i]
		}
		kp += dp[i]
	}
	return np / kp
}
