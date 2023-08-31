package partice

// todo
func findRedundantConnection(edges [][]int) []int {
	res684 := make([]int, 2)
	m := make(map[int]struct{})
	for i, v := range edges {
		_, ok1 := m[v[0]]
		_, ok2 := m[v[1]]
		if ok1 && ok2 {
			if !haveRing(edges[:i-1]) && haveRing(edges[:i]) {
				res684[0] = v[0]
				res684[1] = v[1]
			}
		}
		m[v[0]] = struct{}{}
		m[v[1]] = struct{}{}
	}
	return res684
}

func haveRing(e [][]int) bool {
	edges := make([][]int, len(e))
	copy(edges, e)
	for {
		m := make(map[int]int)
		for _, v := range edges {
			if v == nil {
				continue
			}
			m[v[0]]++
			m[v[1]]++
		}
		end := true
		for i, v := range edges {
			if v == nil {
				continue
			}
			if m[v[0]] <= 1 {
				m[v[1]]--
				edges[i] = nil
				end = false
			}
			if m[v[1]] <= 1 {
				m[v[0]]--
				edges[i] = nil
				end = false
			}
		}
		if len(m) == 0 {
			break
		}
		if end {
			return true
		}
	}
	return false
}
