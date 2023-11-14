package offer

import (
	"fmt"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				l1: &ListNode{Val: 1, Next: &ListNode{Val: 9}},
				l2: &ListNode{Val: 1},
			},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 0}},
		},
		{
			name: "2",
			args: args{
				l1: &ListNode{Val: 9, Next: &ListNode{Val: 9}},
				l2: &ListNode{Val: 1},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.args.l1, tt.args.l2)
			for g := got; g != nil; g = g.Next {
				fmt.Printf("%d", g.Val)
				if g.Next != nil {
					fmt.Print("->")
				} else {
					fmt.Println()
				}
			}
		})
	}
}
