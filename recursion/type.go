package recursion

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(h *ListNode) {
	if h == nil {
		return
	}
	fmt.Printf("%v", h.Val)
	if h.Next != nil {
		fmt.Print("->")
		printList(h.Next)
	} else {
		fmt.Println()
	}
}
