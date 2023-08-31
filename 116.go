package partice

import "fmt"

func connect(root *Node117) *Node117 {
	if root == nil || root.Right == nil {
		return root
	}
	l := &Node117{}
	for h := root; h != nil; h = h.Next {
		h.Left.Next = h.Right
		l.Next = h.Left
		l = h.Right
	}
	for r := root.Left; r != nil; r = r.Next {
		fmt.Println(r.Val)
	}
	connect(root.Left)

	return root
}
