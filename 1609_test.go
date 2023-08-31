package partice

import "testing"

func Test_isEvenOddTree(t *testing.T) {
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
			args: args{
				root: generateTree([]string{`1`, `10`, `4`, `3`, `null`, `7`, `9`, `12`, `8`, `6`, `null`, `null`, `2`}, 0),
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				root: generateTree([]string{`5`, `4`, `2`, `3`, `3`, `7`}, 0),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEvenOddTree(tt.args.root); got != tt.want {
				t.Errorf("isEvenOddTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
