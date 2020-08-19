package partice

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
	m2 := make(map[*Node]int)
	ary := make([]*Node, 0)
	ary2 := make([]*Node, 0)
	result := &Node{}
	p := result
	i := 0
	for ; head != nil; head = head.Next {
		m[head] = -1
		m2[head] = i
		p.Next = &Node{Val: head.Val}
		p = p.Next
		ary = append(ary, p)
		ary2 = append(ary2, head)
		i++
	}
	for k := range m {
		if k.Random == nil {
			continue
		}
		m[k] = m2[k.Random]
	}
	for k := range ary {
		if m[ary2[k]] != -1 {
			ary[k].Random = ary[m[ary2[k]]]
		}
	}

	return result.Next
}

func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}
	l := head
	for l != nil {
		n := &Node{Val: l.Val, Next: l.Next}
		c := l.Next
		l.Next = n
		l = c
	}

	l = head
	for l != nil && l.Next != nil {
		if l.Random != nil {
			l.Next.Random = l.Random.Next
		}
		l = l.Next.Next
	}

	l = head
	r := l.Next
	for l.Next != nil {
		tmp := l.Next
		l.Next = l.Next.Next
		l = tmp
	}
	return r
}
