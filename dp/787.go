package dp

/*
dp[i][j]表示为从src开始，经过i次换成到达j所需的最短消费
*/
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	const inf = 10000*101 + 1
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = inf
		}
	}
	for _, v := range flights {
		if v[0] == src {
			dp[0][v[1]] = v[2]
		}
	}
	for i := 1; i <= k; i++ {
		for _, v := range flights {
			s, d, c := v[0], v[1], v[2]
			if dp[i-1][s] != 0 {
				dp[i][d] = Min(dp[i][d], dp[i-1][s]+c)
			}
		}
	}
	res := inf
	for i := 0; i <= k; i++ {
		if dp[i][dst] != 0 && dp[i][dst] < res {
			res = dp[i][dst]
		}
	}
	if res == inf {
		return -1
	}

	return res
}

func findCheapestPrice2(n int, flights [][]int, src int, dst int, k int) int {
	const inf = 10000*101 + 1
	dp := make([]int, n)
	for j := 0; j < n; j++ {
		dp[j] = inf
	}

	for _, v := range flights {
		if v[0] == src {
			dp[v[1]] = v[2]
		}
	}

	res := dp[dst]
	var tmp []int
	for i := 1; i <= k; i++ {
		tmp = make([]int, n)
		for _, v := range flights {
			s, d, c := v[0], v[1], v[2]
			if dp[s] != 0 && (tmp[d] == 0 || tmp[d] > dp[s]+c) {
				tmp[d] = dp[s] + c
			}
		}
		dp = tmp
		if dp[dst] != 0 && dp[dst] < res {
			res = dp[dst]
		}
	}

	if res == inf {
		return -1
	}
	return res
}
