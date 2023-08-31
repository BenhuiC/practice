package partice

type BSTIterator struct {
	result []int
	idx    int
}

func Constructor173(root *TreeNode) BSTIterator {
	r := make([]int, 0)
	var inorder func(n *TreeNode)
	inorder = func(n *TreeNode) {
		if n == nil {
			return
		}
		if n.Left != nil {
			inorder(n.Left)
		}
		r = append(r, n.Val)
		if n.Right != nil {
			inorder(n.Right)
		}
	}
	inorder(root)
	return BSTIterator{result: r}
}

func (this *BSTIterator) Next() int {
	if len(this.result) == 0 || this.idx > len(this.result) {
		return 0
	}
	idx := this.idx
	this.idx++
	return this.result[idx]
}

func (this *BSTIterator) HasNext() bool {
	return this.idx < len(this.result)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
