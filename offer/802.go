package offer

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	color := make([]int, n)

	var dfs func(i int) bool
	dfs = func(i int) bool {
		if color[i] != 0 {
			return color[i] == 2
		}
		color[i] = 1
		for _, v := range graph[i] {
			if !dfs(v) {
				return false
			}
		}
		color[i] = 2
		return true
	}

	res := make([]int, 0)
	for i := 0; i < n; i++ {
		if dfs(i) {
			res = append(res, i)
		}
	}
	return res
}
