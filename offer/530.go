package offer

import "math"

func getMinimumDifference(root *TreeNode) (res int) {
	res = math.MaxInt
	if root == nil {
		return
	}
	ary := make([]int, 0)
	var inOrder func(n *TreeNode)
	inOrder = func(n *TreeNode) {
		if n == nil {
			return
		}
		inOrder(n.Left)
		ary = append(ary, n.Val)
		inOrder(n.Right)
	}
	inOrder(root)
	for i := 1; i < len(ary); i++ {
		if t := ary[i] - ary[i-1]; t < res {
			res = t
		}
	}

	return res
}
