package tree

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	if root == nil || target == nil {
		return nil
	}
	if k == 0 {
		return []int{target.Val}
	}
	res := make([]int, 0)
	var findChild func(n *TreeNode, l int)
	findChild = func(n *TreeNode, l int) {
		if n == nil || l > k {
			return
		}
		if l == k {
			res = append(res, n.Val)
			return
		}
		findChild(n.Left, l+1)
		findChild(n.Right, l+1)
	}
	findChild(target, 0)

	father := make(map[int]*TreeNode)
	var getFather func(n *TreeNode)
	getFather = func(n *TreeNode) {
		if n == nil {
			return
		}
		if n.Left != nil {
			father[n.Left.Val] = n
			getFather(n.Left)
		}
		if n.Right != nil {
			father[n.Right.Val] = n
			getFather(n.Right)
		}
	}
	getFather(root)

	var findFather func(n *TreeNode, from, l int)
	findFather = func(n *TreeNode, from, l int) {
		if n == nil || l > k {
			return
		}
		if l == k {
			res = append(res, n.Val)
			return
		}
		if n.Left != nil && n.Left.Val == from {
			findChild(n.Right, l+1)
		} else {
			findChild(n.Left, l+1)
		}
		findFather(father[n.Val], n.Val, l+1)
	}
	findFather(father[target.Val], target.Val, 1)

	return res
}
