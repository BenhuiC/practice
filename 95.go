package partice

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generate(1, n)
}

func generate(left, right int) []*TreeNode {
	result := make([]*TreeNode, 0)
	if left > right {
		result = append(result, nil)
		return result
	}
	for i := left; i <= right; i++ {
		left := generate(left, i-1)
		right := generate(i+1, right)
		for _, l := range left {
			for _, r := range right {
				t := TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				}
				result = append(result, &t)
			}
		}
	}

	return result
}
