package partice

import "fmt"

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	if root == nil {
		return res
	}
	var dfs func(n *TreeNode, s string)
	dfs = func(n *TreeNode, s string) {
		if n == nil {
			return
		}
		if n.Left == nil && n.Right == nil {
			s = fmt.Sprintf("%s%d", s, n.Val)
			res = append(res, s)
			return
		}
		s = fmt.Sprintf("%s%d->", s, n.Val)
		dfs(n.Left, s)
		dfs(n.Right, s)
	}
	dfs(root, "")
	return res
}
