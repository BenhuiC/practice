package partice

import "testing"

func Test_longestUnivaluePath(t *testing.T) {
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
				root: generateTree([]string{"5", "4", "5", "1", "1", "5"}, 0),
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				root: generateTree([]string{"1", "4", "5", "4", "4", "5"}, 0),
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				root: generateTree([]string{"1", "2", "2", "2", "2", "2", "2", "2"}, 0),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestUnivaluePath(tt.args.root); got != tt.want {
				t.Errorf("longestUnivaluePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
