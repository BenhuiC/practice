package tree

import "testing"

func Test_smallestFromLeaf(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				root: &TreeNode{Val: 4, Left: &TreeNode{Val: 0, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 1}},
			},
			want: "bae",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestFromLeaf(tt.args.root); got != tt.want {
				t.Errorf("smallestFromLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}
