package offer

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	var dfs func(n *TreeNode, h int)
	dfs = func(n *TreeNode, h int) {
		if n == nil {
			return
		}
		h++
		if h > len(res) {
			res = append(res, n.Val)
		} else if res[h-1] < n.Val {
			res[h-1] = n.Val
		}
		dfs(n.Left, h)
		dfs(n.Right, h)
	}
	dfs(root, 0)
	return res
}
