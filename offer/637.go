package offer

func averageOfLevels(root *TreeNode) []float64 {
	sum := make([]float64, 0)
	count := make([]int, 0)
	var dfs func(n *TreeNode, l int)
	dfs = func(n *TreeNode, l int) {
		if n == nil {
			return
		}
		if l >= len(sum) {
			sum = append(sum, float64(n.Val))
			count = append(count, 1)
		} else {
			sum[l] += float64(n.Val)
			count[l] += 1
		}
		dfs(n.Left, l+1)
		dfs(n.Right, l+1)
	}
	dfs(root, 0)
	res := make([]float64, len(sum))
	for i := 0; i < len(sum); i++ {
		res[i] = sum[i] / float64(count[i])
	}
	return res
}
