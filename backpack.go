package partice

func backPack(val []int, c int) int {
	dp := make([][]int, len(val)+1)
	for i := 0; i <= len(val); i++ {
		dp[i] = make([]int, c+1)
	}
	for y := 1; y <= c; y++ {
		for x := 1; x <= len(val); x++ {
			if y-val[x-1] >= 0 {
				dp[x][y] = Max(dp[x-1][y], dp[x-1][y-val[x-1]]+val[x-1])
			} else {
				dp[x][y] = dp[x-1][y]
			}
		}
	}
	return dp[len(val)][c]
}

func backPack2(val []int, c int) int {
	dp := make([]int, c+1)
	for i := 0; i < len(val); i++ {
		for j := c; j > 0; j-- {
			if j >= val[i] {
				dp[j] = Max(dp[j], dp[j-val[i]]+val[i])
			}
		}
	}
	return dp[c]
}

func backPack3(val, wei []int, c int) int {
	dp := make([][]int, len(val)+1)
	for i := 0; i <= len(val); i++ {
		dp[i] = make([]int, c+1)
	}
	for y := 1; y <= c; y++ {
		for x := 1; x <= len(val); x++ {
			if y >= wei[x-1] {
				dp[x][y] = Max(dp[x-1][y], dp[x-1][y-wei[x-1]]+val[x-1])
			} else {
				dp[x][y] = dp[x-1][y]
			}
		}
	}
	return dp[len(val)][c]
}

func backPack4(val, wei []int, c int) int {
	dp := make([]int, c+1)
	for i := 0; i < len(val); i++ {
		for j := c; j > 0; j-- {
			if j >= wei[i] {
				dp[j] = Max(dp[j], dp[j-wei[i]]+val[i])
			}
		}
	}
	return dp[c]
}
