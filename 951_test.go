package partice

import "testing"

func Test_flipEquiv(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				root1: generateTree([]string{"6", "1", "0"}, 0),
				root2: generateTree([]string{"6", "null", "1"}, 0),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipEquiv(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("flipEquiv() = %v, want %v", got, tt.want)
			}
		})
	}
}
