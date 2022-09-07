package partice

import "testing"

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				root:      generateTree([]string{"5", "4", "8", "11", "null", "13", "4", "7", "2", "null", "null", "null", "1"}, 0),
				targetSum: 22,
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				root:      generateTree([]string{"-2", "null", "-3"}, 0),
				targetSum: -5,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
