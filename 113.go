package partice

func pathSum(root *TreeNode, targetSum int) [][]int {
	res113 := make([][]int, 0)
	ans := make([]int, 0)
	var backTrace func(n *TreeNode)
	backTrace = func(n *TreeNode) {
		if n == nil {
			return
		}
		targetSum -= n.Val
		ans = append(ans, n.Val)
		if targetSum == 0 && n.Left == nil && n.Right == nil {
			res113 = append(res113, append([]int{}, ans...))
			ans = ans[:len(ans)-1]
			targetSum += n.Val
			return
		}
		backTrace(n.Left)
		backTrace(n.Right)
		targetSum += n.Val
		ans = ans[:len(ans)-1]
	}
	backTrace(root)
	return res113
}
