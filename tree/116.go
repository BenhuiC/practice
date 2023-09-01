package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	for h := root; h != nil; h = h.Left {
		var tmp *Node
		for v := h; v != nil; v = v.Next {
			if v.Left != nil {
				if tmp != nil {
					tmp.Next = v.Left
				}
				tmp = v.Left
			}
			if v.Right != nil {
				if tmp != nil {
					tmp.Next = v.Right
				}
				tmp = v.Right
			}
		}
	}

	return root
}
