package tree

func smallestFromLeaf(root *TreeNode) string {
	if root == nil {
		return ""
	}
	res := "~"
	reverseStr := func(str string) string {
		bt := make([]byte, len(str))
		for i := 0; i < len(str); i++ {
			bt[i] = str[len(str)-i-1]
		}
		return string(bt)
	}

	var dfs func(node *TreeNode, bt []byte)
	dfs = func(node *TreeNode, bt []byte) {
		if node == nil {
			return
		}
		bt = append(bt, byte(node.Val+'a'))
		if node.Left == nil && node.Right == nil {
			curStr := reverseStr(string(bt))
			if curStr < res {
				res = curStr
			}
			return
		}
		dfs(node.Left, bt)
		dfs(node.Right, bt)
		return
	}
	dfs(root, []byte{})

	return res
}
