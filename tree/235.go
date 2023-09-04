package tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if root.Val == p.Val || root.Val == q.Val {
			return root
		}
		if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
			continue
		}
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
			continue
		}
		return root
	}

	return nil
}
