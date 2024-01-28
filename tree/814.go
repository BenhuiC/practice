package tree

func pruneTree(root *TreeNode) *TreeNode {
	var ptree func(n *TreeNode) (oneCnt int, t *TreeNode)
	ptree = func(n *TreeNode) (oneCnt int, t *TreeNode) {
		if n == nil {
			return
		}
		lc, ln := ptree(n.Left)
		rc, rn := ptree(n.Right)

		n.Left = ln
		n.Right = rn
		oneCnt = lc + rc
		if n.Val == 1 {
			oneCnt++
		}
		if oneCnt == 0 {
			return 0, nil
		}
		return oneCnt, n
	}

	_, res := ptree(root)
	return res
}
