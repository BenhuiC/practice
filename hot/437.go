package hot

func pathSum(root *TreeNode, targetSum int) int {
	res := 0

	var sums func(n *TreeNode) []int
	sums = func(n *TreeNode) []int {
		if n == nil {
			return nil
		}
		if n.Val == targetSum {
			res++
		}
		left, right := sums(n.Left), sums(n.Right)
		ary := make([]int, 0)
		ary = append(ary, n.Val)
		for _, v := range left {
			tmp := v + n.Val
			ary = append(ary, tmp)
			if tmp == targetSum {
				res++
			}
		}
		for _, v := range right {
			tmp := v + n.Val
			ary = append(ary, tmp)
			if tmp == targetSum {
				res++
			}
		}
		return ary
	}
	sums(root)
	return res
}
