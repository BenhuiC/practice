package hot

func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{}
	h := head

	m := make(minheap, 0)
	for _, l := range lists {
		if l == nil {
			continue
		}
		m.push(l)
	}

	for len(m) > 0 {
		n := m.pop()
		if n == nil {
			break
		}
		if n.Next != nil {
			m.push(n.Next)
		}
		n.Next = nil
		h.Next = n
		h = n
	}

	return head.Next
}

type minheap []*ListNode

func (m *minheap) up() {
	for n := len(*m) - 1; n != 0; n = (n - 1) / 2 {
		r := (n - 1) / 2
		if (*m)[r].Val > (*m)[n].Val {
			(*m)[r], (*m)[n] = (*m)[n], (*m)[r]
			continue
		} else {
			break
		}
	}
}

func (m *minheap) down() {
	i := 0
	n := len(*m)
	for l := i*2 + 1; l < n; l = i*2 + 1 {
		if l < n-1 && (*m)[l+1].Val < (*m)[l].Val {
			l++
		}
		if (*m)[i].Val > (*m)[l].Val {
			(*m)[i], (*m)[l] = (*m)[l], (*m)[i]
			i = l
		} else {
			break
		}
	}
}

func (m *minheap) push(n *ListNode) {
	*m = append(*m, n)
	m.up()
}

func (m *minheap) pop() *ListNode {
	n := len(*m)
	if n == 0 {
		return nil
	}
	res := (*m)[0]
	(*m)[0] = (*m)[n-1]
	*m = (*m)[:n-1]
	m.down()

	return res
}
