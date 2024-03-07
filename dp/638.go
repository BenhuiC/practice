package dp

func shoppingOffers(price []int, special [][]int, needs []int) int {
	n := len(needs)
	if n <= 0 {
		return 0
	}
	usefulSpecial := make([][]int, 0)
	for i, s := range special {
		cost := 0
		add := true
		for j, p := range s[:n] {
			if p > needs[j] {
				add = false
				break
			}
			cost += p * price[j]
		}
		if add && cost > s[n] {
			usefulSpecial = append(usefulSpecial, special[i])
		}
	}

	dp := map[string]int{}
	var dfs func(need []byte) int
	dfs = func(need []byte) int {
		if v, ok := dp[string(need)]; ok {
			return v
		}
		minPrice := 0
		for i, p := range need {
			minPrice += int(p) * price[i]
		}

		nextNeed := make([]byte, n)
		for _, s := range usefulSpecial {
			canUse := true
			for i, p := range s[:n] {
				if p > int(need[i]) {
					canUse = false
					break
				}
				nextNeed[i] = need[i] - byte(p)
			}
			if canUse {
				minPrice = min(minPrice, dfs(nextNeed)+s[n])
			}
		}

		dp[string(need)] = minPrice
		return minPrice
	}

	needByte := make([]byte, n)
	for i, v := range needs {
		needByte[i] = byte(v)
	}

	return dfs(needByte)
}
