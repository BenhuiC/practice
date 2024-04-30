package tree

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	n := len(voyage)
	if n == 0 {
		return []int{-1}
	}

	res := make([]int, 0)
	idx := 0
	var flipMatchIdx func(node *TreeNode) bool
	flipMatchIdx = func(node *TreeNode) bool {
		if node == nil && idx == n {
			return true
		}
		if node.Val != voyage[idx] {
			return false
		}
		idx++
		if node.Left == nil && node.Right == nil {
			return true
		}

		if node.Left != nil && node.Right != nil {
			if node.Left.Val == voyage[idx] {
				l := flipMatchIdx(node.Left)
				if !l {
					return false
				}
				return flipMatchIdx(node.Right)
			} else {
				res = append(res, node.Val)
				l := flipMatchIdx(node.Right)
				if !l {
					return false
				}
				return flipMatchIdx(node.Left)
			}
		} else if node.Left != nil {
			return flipMatchIdx(node.Left)
		} else if node.Right != nil {
			return flipMatchIdx(node.Right)
		}

		return false
	}
	if !flipMatchIdx(root) {
		return []int{-1}
	}

	return res
}
