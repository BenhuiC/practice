package hot

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	left := inorderTraversal(root.Left)
	left = append(left, root.Val)
	left = append(left, inorderTraversal(root.Right)...)
	return left
}

func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, n.Val)
		cur = n.Right
	}

	return res
}
