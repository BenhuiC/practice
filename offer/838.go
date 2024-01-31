package offer

func pushDominoes(dominoes string) string {
	n := len(dominoes)
	dis := make([]int, n)
	res := make([]byte, n)
	ldis := 0
	for i := n - 1; i >= 0; i-- {
		if dominoes[i] == 'L' {
			ldis = 1
			dis[i] = 1
			res[i] = 'L'
			continue
		}
		if dominoes[i] == 'R' {
			ldis = 0
			res[i] = 'R'
			continue
		}
		if dominoes[i] == '.' && ldis > 0 {
			ldis++
			dis[i] = ldis
			res[i] = 'L'
		}
	}

	rdis := 0
	for i := 0; i < n; i++ {
		if dominoes[i] == 'R' {
			rdis = 1
			continue
		}
		if dominoes[i] == 'L' {
			rdis = 0
			continue
		}
		if rdis > 0 {
			rdis++
			if rdis < dis[i] || dis[i] == 0 {
				res[i] = 'R'
			} else if rdis == dis[i] {
				res[i] = '.'
			}
		} else if dis[i] == 0 {
			res[i] = '.'
		}

	}

	return string(res)
}
