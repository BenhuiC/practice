package partice

var vec [][]int

func levelOrder(root *TreeNode) [][]int {
	vec = make([][]int, 0)
	vertical(root, 0)
	return vec
}

func vertical(node *TreeNode, n int) {
	if node == nil {
		return
	}
	if n >= len(vec) {
		vec = append(vec, make([]int, 0))
	}
	vec[n] = append(vec[n], node.Val)
	vertical(node.Left, n+1)
	vertical(node.Right, n+1)
}

func levelOrder2(root *TreeNode) [][]int {
	vec := make([][]int, 0)
	if root == nil {
		return vec
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		l := make([]int, 0)
		for _, v := range queue {
			l = append(l, v.Val)
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		queue = queue[n:]
		vec = append(vec, l)
	}

	return vec
}
