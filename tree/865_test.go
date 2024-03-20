package tree

import (
	"reflect"
	"testing"
)

func Test_subtreeWithAllDeepest(t *testing.T) {
	n0 := &TreeNode{Val: 0}
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
	n6 := &TreeNode{Val: 6}
	n7 := &TreeNode{Val: 7}
	n8 := &TreeNode{Val: 8}
	n9 := &TreeNode{Val: 9}

	n3.Left = n5
	n3.Right = n1
	n5.Left = n6
	n5.Right = n2
	n2.Left = n7
	n2.Right = n4
	n1.Left = n0
	n1.Right = n8

	n0.Left = n9

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "1",
			args: args{
				root: n3,
			},
			want: n3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtreeWithAllDeepest(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subtreeWithAllDeepest() = %v, want %v", got, tt.want)
			}
		})
	}
}
