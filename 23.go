package partice

func mergeKLists(lists []*ListNode) *ListNode {
	h := &ListNode{}
	heap := NewMinHeap()
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Insert(lists[i])
		}
	}
	t := h
	for len(heap) > 0 {
		tmp := heap[0]
		t.Next = tmp
		t = tmp
		heap.DelMin()
		if tmp.Next != nil {
			heap.Insert(tmp.Next)
		}
	}

	return h.Next
}
