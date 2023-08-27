package link

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	}
	fmt.Println(isPalindrome(l))
}
