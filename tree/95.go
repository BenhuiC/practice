package tree

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	var g func(l, r int) []*TreeNode
	g = func(l, r int) []*TreeNode {
		if l > r {
			return []*TreeNode{nil} // 这里必须要传nil，下面三层循环的第二次会先遍历左节点
		}
		nodes := make([]*TreeNode, 0)
		for cur := l; cur <= r; cur++ {
			lnodes := g(l, cur-1)
			rnodes := g(cur+1, r)
			for _, ln := range lnodes {
				for _, rn := range rnodes {
					node := &TreeNode{
						Val:   cur,
						Left:  ln,
						Right: rn,
					}
					nodes = append(nodes, node)
				}
			}
		}
		return nodes
	}

	return g(1, n)
}
