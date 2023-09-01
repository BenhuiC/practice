package tree

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	res := make([]int, 0)
	for len(stack) > 0 {
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, t.Val)
		if t.Right != nil {
			stack = append(stack, t.Right)
		}
		if t.Left != nil {
			stack = append(stack, t.Left)
		}
	}
	return res
}
