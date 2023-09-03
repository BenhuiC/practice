package tree

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	res := make([]*TreeNode, 0)

	type pair struct {
		i int
		n *TreeNode
	}
	repeat := make(map[[3]int]pair)
	resMap := make(map[*TreeNode]struct{})
	idx := 0
	var dfs func(n *TreeNode) int
	dfs = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		p := [3]int{n.Val, dfs(n.Left), dfs(n.Right)}
		if v, ok := repeat[p]; ok {
			resMap[v.n] = struct{}{}
			return v.i
		}
		idx++
		repeat[p] = pair{
			i: idx,
			n: n,
		}
		return idx
	}
	dfs(root)
	for k := range resMap {
		res = append(res, k)
	}

	return res
}
