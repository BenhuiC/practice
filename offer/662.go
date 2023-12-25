package offer

func widthOfBinaryTree(root *TreeNode) int {
	res := 0
	ary := make([][2]int, 0)
	var dfs func(n *TreeNode, d, x int)
	dfs = func(n *TreeNode, d, x int) {
		if n == nil {
			return
		}
		if d >= len(ary) {
			ary = append(ary, [2]int{x, x})
		}
		if x < ary[d][0] {
			ary[d][0] = x
		}
		if x > ary[d][1] {
			ary[d][1] = x
		}
		if ary[d][1]-ary[d][0]+1 > res {
			res = ary[d][1] - ary[d][0] + 1
		}
		dfs(n.Left, d+1, 2*x-1)
		dfs(n.Right, d+1, 2*x)
	}
	dfs(root, 0, 1)
	return res
}
