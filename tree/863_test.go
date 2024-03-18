package tree

import (
	"reflect"
	"testing"
)

func Test_distanceK(t *testing.T) {
	n0 := &TreeNode{Val: 0}
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
	n6 := &TreeNode{Val: 6}

	n0.Left = n3
	n0.Right = n1
	n1.Left = n2
	n1.Right = n5
	n3.Left = n4
	n4.Right = n6

	type args struct {
		root   *TreeNode
		target *TreeNode
		k      int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				root:   n0,
				target: n3,
				k:      2,
			},
			want: []int{6, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceK(tt.args.root, tt.args.target, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distanceK() = %v, want %v", got, tt.want)
			}
		})
	}
}
