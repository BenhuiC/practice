package offer

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parents := make([]int, n+1) // parent 记录当前节点的祖先节点
	for i := 0; i <= n; i++ {
		parents[i] = i
	}
	var findAncestor func(x int) int
	findAncestor = func(x int) int {
		if parents[x] != x {
			parents[x] = findAncestor(parents[x])
		}
		return parents[x]
	}
	union := func(from, to int) bool {
		x, y := findAncestor(from), findAncestor(to)
		if x == y { // 如果当前两个节点的祖先节点是同一个，则这两个节点在同一个树上，如果再添加一条边，则会形成环
			return false
		}
		parents[x] = y
		return true
	}
	for _, e := range edges {
		if !union(e[0], e[1]) {
			return e
		}
	}

	return nil
}
