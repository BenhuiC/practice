package link

import (
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	l = reverseBetween(l, 2, 5)
	printList(l)
}
