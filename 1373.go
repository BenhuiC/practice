package partice

var res1373 int

func maxSumBST(root *TreeNode) int {
	res1373 = 0
	_, _, _, _ = sumBST(root)
	if res1373 < 0 {
		res1373 = 0
	}
	return res1373
}

func sumBST(n *TreeNode) (isBST bool, sum, mx, mi int) {
	if n == nil {
		isBST = true
		return
	}
	if n.Right == nil && n.Left == nil {
		isBST = true
		sum, mx, mi = n.Val, n.Val, n.Val
		return
	}
	leftFlag, leftSum, leftMax, leftMin := sumBST(n.Left)
	rightFlag, rightSum, rightMax, rightMin := sumBST(n.Right)
	if n.Right == nil {
		if leftFlag && leftMax < n.Val {
			isBST = true
			sum = n.Val + leftSum
			mi = leftMin
			mx = n.Val
			res1373 = max(res1373, sum)
			return
		}
		sum = leftSum
		res1373 = max(res1373, sum)
		return
	}
	if n.Left == nil {
		if rightFlag && rightMin > n.Val {
			isBST = true
			sum = n.Val + rightSum
			mi = n.Val
			mx = rightMax
			res1373 = max(res1373, sum)
			return
		}
		sum = rightSum
		res1373 = max(res1373, sum)
		return
	}
	if leftFlag && leftMax < n.Val && rightFlag && rightMin > n.Val {
		sum = n.Val + leftSum + rightSum
		mi = leftMin
		mx = rightMax
		isBST = true
		res1373 = max(res1373, sum)
		return
	}
	sum = max(leftSum, rightSum)
	res1373 = max(res1373, sum)
	return
}
