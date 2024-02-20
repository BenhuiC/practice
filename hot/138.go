package hot

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	m := make(map[*Node]int)
	ary := make([]*Node, 0)
	idx := 0
	for h := head; h != nil; h = h.Next {
		ary = append(ary, &Node{Val: h.Val})
		m[h] = idx
		idx++
	}

	i := 0
	for h := head; h != nil; h = h.Next {
		if i != len(ary)-1 {
			ary[i].Next = ary[i+1]
		}
		if h.Random != nil {
			ary[i].Random = ary[m[h.Random]]
		}
		i++
	}

	return ary[0]
}
