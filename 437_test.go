package partice

import "testing"

func Test_pathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				root: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val: 3,
							},
							Right: &TreeNode{
								Val: -2,
							},
						},
						Right: &TreeNode{
							Val: 0,
							Right: &TreeNode{
								Val: 8,
							},
						},
					},
					Right: &TreeNode{
						Val: -3,
						Right: &TreeNode{
							Val: 11,
						},
					},
				},
				targetSum: 8,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
