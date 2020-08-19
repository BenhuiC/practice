package partice

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{root: generateTree([]string{`3`, `9`, `20`, `null`, `null`, `15`, `7`}, 0)},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
