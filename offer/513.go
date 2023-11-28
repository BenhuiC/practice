package offer

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	cur := []*TreeNode{root}
	next := make([]*TreeNode, 0)
	for {
		for _, n := range cur {
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		if len(next) == 0 {
			return cur[0].Val
		}
		cur, next = next, []*TreeNode{}
	}
}

func findBottomLeftValue2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var maxHeight, res int
	var dfs func(n *TreeNode, h int)
	dfs = func(n *TreeNode, h int) {
		if n == nil {
			return
		}
		h++
		if h > maxHeight {
			maxHeight = h
			res = n.Val
		}
		dfs(n.Left, h)
		dfs(n.Right, h)
	}
	dfs(root, 0)
	return res
}
