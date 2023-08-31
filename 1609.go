package partice

func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Val%2 == 0 {
		return false
	}
	cur, next := make([]*TreeNode, 0), make([]*TreeNode, 0)
	cur = append(cur, root)
	level := 1
	for len(cur) != 0 {
		for i := range cur {
			if cur[i].Left != nil {
				next = append(next, cur[i].Left)
			}
			if cur[i].Right != nil {
				next = append(next, cur[i].Right)
			}
		}
		if !checkStd(level, next) {
			return false
		}
		level++
		cur = next
		next = make([]*TreeNode, 0)
	}

	return true
}

func checkStd(lv int, ary []*TreeNode) bool {
	desc := lv%2 == 1
	yu := (lv + 1) % 2
	if len(ary) == 0 {
		return true
	}
	if len(ary) == 1 {
		return ary[0].Val%2 == yu
	}
	for i := 1; i < len(ary); i++ {
		if ary[i].Val%2 != yu {
			return false
		}
		if desc {
			if ary[i].Val < ary[i-1].Val {
				continue
			} else {
				return false
			}
		} else {
			if ary[i].Val > ary[i-1].Val {
				continue
			} else {
				return false
			}
		}
	}
	return true
}
