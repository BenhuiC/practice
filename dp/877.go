package dp

func stoneGame(piles []int) bool {
	n := len(piles)
	if n < 2 {
		return true
	}
	dp1 := make([]int, n)
	dp2 := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		dp1[i] = piles[i]
		for j := i + 1; j < n; j++ {
			left := piles[i] + dp2[j]
			right := piles[j] + dp2[j-1]
			if left >= right {
				dp2[j] = dp1[j]
				dp1[j] = left
			} else {
				dp1[j] = right
				dp2[j] = dp1[j-1]
			}
		}
	}

	return dp1[n-1] > dp2[n-1]
}
