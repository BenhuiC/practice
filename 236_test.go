package partice

import (
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		//{
		//	name: "1",
		//	args: args{root: generateTree([]string{`-1`, `0`, `3`, `-2`, `4`, `null`, `null`, `8`}, 0), p: &TreeNode{Val: 3}, q: &TreeNode{Val: 8}},
		//	want: &TreeNode{Val: -1},
		//},
		{
			name: "2",
			args: args{root: generateTree([]string{`3`, `5`, `1`, `6`, `2`, `0`, `8`, `null`, `null`, `7`, `4`}, 0), p: &TreeNode{Val: 5}, q: &TreeNode{Val: 8}},
			want: &TreeNode{Val: 3},
		},
		{
			name: "3",
			args: args{root: generateTree([]string{`3`, `5`, `1`, `6`, `2`, `0`, `8`, `null`, `null`, `7`, `4`}, 0), p: &TreeNode{Val: 4}, q: &TreeNode{Val: 8}},
			want: &TreeNode{Val: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); got.Val != tt.want.Val {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got.Val, tt.want.Val)
			}
		})
	}
}
