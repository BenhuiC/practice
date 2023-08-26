package link

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	h := &ListNode{}
	e := h
	m := minheap{}
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}
		m.push(lists[i])
	}
	for {
		v := m.pop()
		if v == nil {
			break
		}
		e.Next = v
		e = v
		if v.Next != nil {
			m.push(v.Next)
		}
		e.Next = nil
	}

	return h.Next
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
