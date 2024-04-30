package tree

import (
	"reflect"
	"testing"
)

func Test_flipMatchVoyage(t *testing.T) {
	type args struct {
		root   *TreeNode
		voyage []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				root:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
				voyage: []int{1, 3, 2},
			},
			want: []int{1},
		},
		{
			name: "2",
			args: args{
				root:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}},
				voyage: []int{1, 3, 4, 5, 2},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipMatchVoyage(tt.args.root, tt.args.voyage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flipMatchVoyage() = %v, want %v", got, tt.want)
			}
		})
	}
}
