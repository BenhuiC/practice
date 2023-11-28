package offer

func findFrequentTreeSum(root *TreeNode) []int {
	sumMap := make(map[int]int)
	var treeSum func(n *TreeNode) int
	treeSum = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		left := treeSum(n.Left)
		right := treeSum(n.Right)
		s := n.Val + left + right
		sumMap[s]++
		return s
	}
	_ = treeSum(root)
	var cnt int
	res := make([]int, 0)
	for k, v := range sumMap {
		if v == cnt {
			res = append(res, k)
		} else if v > cnt {
			res = []int{k}
			cnt = v
		}
	}

	return res
}
