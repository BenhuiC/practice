package back_trace

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := minDepth(root.Left)
	r := minDepth(root.Right)
	if root.Left == nil {
		return r + 1
	}
	if root.Right == nil {
		return l + 1
	}
	res := l
	if r < l {
		res = r
	}
	return res + 1
}
