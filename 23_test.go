package partice

import (
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		//{
		//	name: "1",
		//	args: args{
		//		lists: []*ListNode{{Val: -8, Next: &ListNode{Val: -7, Next: &ListNode{Val: -7, Next: &ListNode{Val: -5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}}}}},
		//			{Val: -2},
		//			{Val: -10, Next: &ListNode{Val: -10, Next: &ListNode{Val: -7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}}}},
		//			{Val: 2}},
		//	},
		//	want: nil,
		//},
		{
			name: "2",
			args: args{
				lists: []*ListNode{{Val: -10, Next: &ListNode{Val: -9, Next: &ListNode{Val: -9, Next: &ListNode{Val: -3, Next: &ListNode{Val: -1, Next: &ListNode{Val: -1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}}}}},
					{Val: -5},
					{Val: -9, Next: &ListNode{Val: -6, Next: &ListNode{Val: -5, Next: &ListNode{Val: -4, Next: &ListNode{Val: -2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}}}},
					{Val: 4},
					{Val: -8},
					{Val: -3, Next: &ListNode{Val: -3, Next: &ListNode{Val: -2, Next: &ListNode{Val: -1, Next: &ListNode{Val: 0}}}}},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeKLists(tt.args.lists)
			got.Print()
		})
	}
}
