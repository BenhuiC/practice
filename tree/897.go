package tree

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	nodes := make([]*TreeNode, 0)
	var inorder func(n *TreeNode)
	inorder = func(n *TreeNode) {
		if n == nil {
			return
		}
		inorder(n.Left)
		n.Left = nil
		nodes = append(nodes, n)
		inorder(n.Right)
		n.Right = nil
	}
	inorder(root)

	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Left = nil
		nodes[i].Right = nodes[i+1]
	}

	return nodes[0]
}
