package partice

import (
	"strconv"
)

func printTree(root *TreeNode) [][]string {
	height := treeHeight(root) - 1
	m := height + 1
	n := 1<<m - 1
	ans := make([][]string, m)
	for i := range ans {
		ans[i] = make([]string, n)
	}
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, r, c int) {
		ans[r][c] = strconv.Itoa(node.Val)
		if node.Left != nil {
			dfs(node.Left, r+1, c-1<<(height-r-1))
		}
		if node.Right != nil {
			dfs(node.Right, r+1, c+1<<(height-r-1))
		}
	}
	dfs(root, 0, (n-1)/2)
	return ans
}

func treeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := treeHeight(root.Left)
	right := treeHeight(root.Right)
	return Max(left, right) + 1
}

//输入：root = [1,2,3,null,4]
//输出：
//[["","","","1","","",""],
//["","2","","","","3",""],
//["","","4","","","",""]]
