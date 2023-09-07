package dp

import "testing"

func Test_minFallingPathSum2(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				matrix: [][]int{{100, -42, -46, -41}, {31, 97, 10, -10}, {-58, -51, 82, 89}, {51, 81, 69, -51}},
			},
			want: -36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFallingPathSum2(tt.args.matrix); got != tt.want {
				t.Errorf("minFallingPathSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
