package partice

import (
	"fmt"
	"testing"
)

func Test_removeZeroSumSublists(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: -3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 1}}}}},
			},
		},
		{
			name: "2",
			args: args{
				&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: -3, Next: &ListNode{Val: 4}}}}},
			},
		},
		{
			name: "3",
			args: args{
				&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: -3, Next: &ListNode{Val: -2}}}}},
			},
		},
		{
			name: "4",
			args: args{
				&ListNode{Val: 1, Next: &ListNode{Val: -1}},
			},
		},
		{
			name: "5",
			args: args{
				&ListNode{Val: 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeZeroSumSublists(tt.args.head)
			for g := got; g != nil; g = g.Next {
				fmt.Printf("%v\t", g.Val)
			}
			fmt.Println()
		})
	}
}
