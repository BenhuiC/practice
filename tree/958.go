package tree

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	cur := []*TreeNode{root}
	leaf := false
	for len(cur) != 0 && !leaf {
		next := make([]*TreeNode, 0)
		for _, n := range cur {
			if leaf {
				if n.Left != nil || n.Right != nil {
					return false
				}
			} else {
				if n.Left != nil {
					next = append(next, n.Left)
				} else {
					leaf = true
				}
				if n.Right != nil {
					next = append(next, n.Right)
					if leaf {
						return false
					}
				} else {
					leaf = true
				}
			}
		}
		cur = next
	}

	for _, n := range cur {
		if n.Left != nil || n.Right != nil {
			return false
		}
	}

	return true
}
