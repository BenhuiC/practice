package partice

func allPossibleFBT(n int) []*TreeNode {
	var dfs func(n int) []*TreeNode
	dfs = func(n int) []*TreeNode {
		if n == 0 {
			return nil
		}
		r := make([]*TreeNode, 0)
		if n == 1 {
			r = append(r, &TreeNode{})
			return r
		}
		left := 1
		right := n - 2
		for right >= 1 {
			leftAry := dfs(left)
			rightAry := dfs(right)
			for i := range leftAry {
				for j := range rightAry {
					r = append(r, &TreeNode{Left: leftAry[i], Right: rightAry[j]})
				}
			}

			left += 2
			right -= 2
		}
		return r
	}

	return dfs(n)
}
