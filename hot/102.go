package hot

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	cur := []*TreeNode{root}
	for len(cur) != 0 {
		ary := make([]int, 0)
		next := make([]*TreeNode, 0)
		for _, t := range cur {
			ary = append(ary, t.Val)
			if t.Left != nil {
				next = append(next, t.Left)
			}
			if t.Right != nil {
				next = append(next, t.Right)
			}
		}
		cur = next
		res = append(res, ary)
	}
	return res
}
