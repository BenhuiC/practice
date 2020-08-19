package partice

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	b := false
	for len(queue) > 0 {
		n := len(queue)
		l := make([]int, 0)

		for _, v := range queue {
			if !b {
				l = append(l, v.Val)
			} else {
				tmp := []int{v.Val}
				l = append(tmp, l...)
			}

			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}

		queue = queue[n:]
		result = append(result, l)
		b = !b
	}

	return result
}

var vecs [][]int

func zigzagLevelOrder2(root *TreeNode) [][]int {
	zig(root, 0)
	return vecs
}

func zig(n *TreeNode, height int) {
	if n == nil {
		return
	}
	if height >= len(vecs) {
		vecs = append(vecs, make([]int, 0))
	}
	if height%2 == 0 {
		vecs[height] = append(vecs[height], n.Val)
	} else {
		tmp := []int{n.Val}
		vecs[height] = append(tmp, vecs[height]...)
	}
	zig(n.Left, height+1)
	zig(n.Right, height+1)
}
