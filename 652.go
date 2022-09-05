package partice

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res := make([]*TreeNode, 0)
	type pair struct {
		idx  int
		node *TreeNode
	}
	repeat := map[[3]int]pair{}
	resMap := map[*TreeNode]struct{}{}
	idx := 0
	var dfs func(t *TreeNode) int
	dfs = func(t *TreeNode) int {
		if t == nil {
			return 0
		}
		tir := [3]int{t.Val, dfs(t.Left), dfs(t.Right)}
		if p, ok := repeat[tir]; ok {
			resMap[p.node] = struct{}{}
			return p.idx
		}
		idx++
		repeat[tir] = pair{
			idx:  idx,
			node: t,
		}
		return idx
	}
	dfs(root)

	for k := range resMap {
		res = append(res, k)
	}

	return res
}
