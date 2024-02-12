package hot

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			},
			want: []int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal2() = %v, want %v", got, tt.want)
			}
		})
	}
}
