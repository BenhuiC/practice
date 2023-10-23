package offer

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return nil
	}

	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	parents := make(map[int]int)
	var bfs = func(start int) int {
		visit := make([]bool, n)
		visit[start] = true
		q := []int{start}
		var x int
		for len(q) > 0 {
			x, q = q[0], q[1:]
			for _, y := range g[x] {
				if !visit[y] {
					q = append(q, y)
					parents[y] = x
					visit[y] = true
				}
			}
		}

		return x
	}

	p1 := bfs(0)
	p2 := bfs(p1)

	path := make([]int, 0)
	parents[p1] = -1
	for p2 != -1 {
		path = append(path, p2)
		p2 = parents[p2]
	}

	m := len(path)
	if m&1 == 0 {
		return []int{path[m/2-1], path[m/2]}
	}

	return []int{path[m/2]}
}
