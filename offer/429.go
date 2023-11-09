package offer

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	ary := make([]*Node, 0, 1)
	res := make([][]int, 0)
	ary = append(ary, root)
	for len(ary) != 0 {
		n := len(ary)
		l := make([]int, 0, n)
		for i := 0; i < n; i++ {
			l = append(l, ary[i].Val)
			if len(ary[i].Children) == 0 {
				continue
			}
			for t := range ary[i].Children {
				ary = append(ary, ary[i].Children[t])
			}
		}
		res = append(res, l)
		ary = ary[n:]
	}

	return res
}
