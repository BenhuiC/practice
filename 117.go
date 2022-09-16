package partice

type Node117 struct {
	Val   int
	Left  *Node117
	Right *Node117
	Next  *Node117
}

func connect117(root *Node117) *Node117 {
	if root == nil {
		return root
	}
	l := &Node117{}
	p := l
	for h := root; h != nil; h = h.Next {
		if h.Left != nil {
			l.Next = h.Left
			l = l.Next
		}
		if h.Right != nil {
			l.Next = h.Right
			l = l.Next
		}
	}
	//for r := root; r != nil; r = r.Next {
	//	fmt.Printf("%d\t", r.Val)
	//}
	//fmt.Println()
	connect117(p.Next)

	return root
}
