package partice

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, root)
	for len(nodes) != 0 {
		n := len(nodes)
		res = append(res, nodes[n-1].Val)
		for i := 0; i < n; i++ {
			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
			}
		}
		nodes = nodes[n:]
	}

	return res
}

func rightSideViewWithDep(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	var dep func(n *TreeNode, depth int)
	dep = func(n *TreeNode, depth int) {
		if n == nil {
			return
		}
		res = append(res, n.Val)
		dep(n.Right, depth+1)

	}

	return res
}
