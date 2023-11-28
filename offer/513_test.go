package offer

import "testing"

func Test_findBottomLeftValue2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				root: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBottomLeftValue2(tt.args.root); got != tt.want {
				t.Errorf("findBottomLeftValue2() = %v, want %v", got, tt.want)
			}
		})
	}
}
