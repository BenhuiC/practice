package partice

import (
	"reflect"
	"testing"
)

func Test_delete(t *testing.T) {
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
			args: args{root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 4,
				},
			}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteN(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
