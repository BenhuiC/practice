package hot

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	cur := []*TreeNode{root}
	for len(cur) > 0 {
		res = append(res, cur[len(cur)-1].Val)
		next := make([]*TreeNode, 0)
		for _, c := range cur {
			if c.Left != nil {
				next = append(next, c.Left)
			}
			if c.Right != nil {
				next = append(next, c.Right)
			}
		}
		cur = next
	}

	return res
}
