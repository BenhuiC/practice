package partice

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{root: generateTree([]string{"2", "1", "3"}, 0)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST2(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
