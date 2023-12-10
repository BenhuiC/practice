package offer

func leastBricks(wall [][]int) int {
	m := len(wall)
	if m == 0 {
		return 0
	}
	mp := make(map[int]int)
	for i := 0; i < m; i++ {
		lineSum := 0
		for j := 0; j < len(wall[i])-1; j++ {
			lineSum += wall[i][j]
			mp[lineSum]++
		}
	}

	res := 0
	for _, v := range mp {
		if v > res {
			res = v
		}
	}
	res = m - res

	return res
}
