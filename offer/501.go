package offer

func findMode(root *TreeNode) []int {
	res := make([]int, 0)
	var base, count, maxCount int
	var update = func(val int) {
		if val == base {
			count++
		} else {
			base, count = val, 1
		}
		if count == maxCount {
			res = append(res, val)
		} else if count > maxCount {
			maxCount = count
			res = []int{val}
		}
	}
	cur := root
	for cur != nil {
		if cur.Left == nil {
			update(cur.Val)
			cur = cur.Right
			continue
		}
		pre := cur.Left
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}
		if pre.Right == nil {
			pre.Right = cur
			cur = cur.Left
		} else {
			pre.Right = nil
			update(cur.Val)
			cur = cur.Right
		}
	}
	return res
}
