package offer

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	for n := root; n != nil; {
		if n.Child == nil {
			n = n.Next
			continue
		}
		child := flatten(n.Child)
		next := n.Next
		n.Next = child
		child.Prev = n
		childLast := child
		for childLast.Next != nil {
			childLast = childLast.Next
		}
		childLast.Next = next
		if next != nil {
			next.Prev = childLast
		}
		n.Child = nil
		n = next
	}
	return root
}
