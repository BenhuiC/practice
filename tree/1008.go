package tree

func bstFromPreorder(preorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	mid := findMid(preorder, 1, n, preorder[0])
	root.Left = bstFromPreorder(preorder[1:mid])
	root.Right = bstFromPreorder(preorder[mid:])
	return root
}

func findMid(ary []int, left, right, target int) int {
	for left < right {
		mid := left + (right-left)/2
		if ary[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
