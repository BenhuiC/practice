package offer

func knightProbability(n int, k int, row int, column int) float64 {
	direction := [][2]int{{-2, -1}, {-2, 1}, {-1, -2}, {1, -2}, {2, -1}, {2, 1}, {1, 2}, {-1, 2}}
	dp := make([][][]float64, k+1)
	for step := 0; step <= k; step++ {
		dp[step] = make([][]float64, n)
		for i := 0; i < n; i++ {
			dp[step][i] = make([]float64, n)
			for j := 0; j < n; j++ {
				if step == 0 {
					dp[step][i][j] = 1
					continue
				}
				for _, d := range direction {
					ti, tj := i+d[0], j+d[1]
					if ti >= 0 && ti < n && tj >= 0 && tj < n {
						dp[step][i][j] += 0.125 * dp[step-1][ti][tj]
					}
				}
			}
		}
	}
	return dp[k][row][column]
}
