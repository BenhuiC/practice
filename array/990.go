package array

func equationsPossible(equations []string) bool {
	parents := make([]int, 26)
	for i := 0; i < 26; i++ {
		parents[i] = i
	}

	find := func(idx int) int {
		for parents[idx] != idx {
			parents[idx] = parents[parents[idx]]
			idx = parents[idx]
		}
		return parents[idx]
	}
	union := func(i, j int) {
		parents[find(i)] = find(j)
	}

	for _, v := range equations {
		if v[1] == '=' {
			idx1, idx2 := v[0]-'a', v[3]-'a'
			union(int(idx1), int(idx2))
		}
	}
	for _, v := range equations {
		if v[1] == '!' {
			idx1, idx2 := v[0]-'a', v[3]-'a'
			if find(int(idx1)) == find(int(idx2)) {
				return false
			}
		}
	}
	return true
}
