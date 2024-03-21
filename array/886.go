package array

func possibleBipartition(n int, dislikes [][]int) bool {
	dis := make([][]int, n)
	for _, d := range dislikes {
		a, b := d[0]-1, d[1]-1
		dis[a] = append(dis[a], b)
		dis[b] = append(dis[b], a)
	}

	ary := make([]int, n)
	var dfs func(idx, c int) bool
	dfs = func(idx, c int) bool {
		ary[idx] = c
		for _, d := range dis[idx] {
			if ary[d] == -c {
				continue
			}
			if ary[d] == c {
				return false
			}
			if !dfs(d, -c) {
				return false
			}
		}
		return true
	}

	for i := 0; i < n; i++ {
		if ary[i] != 0 {
			continue
		}
		if !dfs(i, 1) {
			return false
		}
	}

	return true
}
