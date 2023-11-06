package offer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	var solu func(n *TreeNode, left bool) int
	solu = func(n *TreeNode, left bool) int {
		if n == nil {
			return 0
		}
		if n.Left == nil && n.Right == nil {
			if left {
				return n.Val
			}
			return 0
		}
		return solu(n.Left, true) + solu(n.Right, false)
	}
	return solu(root, false)
}
