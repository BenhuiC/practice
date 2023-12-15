package offer

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	head := &TreeNode{Left: root}
	var addOne func(n *TreeNode, d int) *TreeNode
	addOne = func(n *TreeNode, d int) *TreeNode {
		if n == nil {
			return n
		}
		if d > 1 {
			n.Left = addOne(n.Left, d-1)
			n.Right = addOne(n.Right, d-1)
			return n
		}
		l := &TreeNode{Val: val, Left: n.Left}
		r := &TreeNode{Val: val, Right: n.Right}
		n.Left = l
		n.Right = r
		return n
	}
	head = addOne(head, depth)

	return head.Left
}
