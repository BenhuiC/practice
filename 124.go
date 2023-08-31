package partice

func maxPathSum(root *TreeNode) int {
	mm := ^int(^uint(0) >> 1)
	_ = maxSum(root, &mm)
	return mm
}

func maxSum(root *TreeNode, n *int) int {
	var m int
	if root == nil {
		return m
	}
	left := maxSum(root.Left, n)
	right := maxSum(root.Right, n)

	m = Max(Max(left+root.Val, right+root.Val), Max(root.Val, left+right+root.Val))
	if m > *n {
		*n = m
	}

	return Max(Max(left+root.Val, right+root.Val), root.Val)
}

/*
 1,2,3,4,  5  ,6 ,7,8,9,10  ,11  ,12  ,13
[5,4,8,11,null,13,4,7,2,null,null,null,1]
		 5
	  4      8
	11 nil  13  4
  7  2        1

*/
