package partice

import "testing"

func Test_countSquares(t *testing.T) {
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
				matrix: [][]int{{1,0,1,0,0},{1,0,1,1,1},{1,1,1,1,1},{1,0,0,1,0}},
			},
			want: 15,
		},
		{
			name: "2",
			args: args{
				matrix: [][]int{{1,1,1,1,0},{1,1,1,1,0},{1,1,1,1,1},{1,1,1,1,1},{0,0,1,1,1}},
			},
			want: 39,
		},
		{
			name: "3",
			args: args{
				matrix: [][]int{{0,0,0,1},{1,1,0,1},{1,1,1,1},{0,1,1,1},{0,1,1,1}},
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSquares(tt.args.matrix); got != tt.want {
				t.Errorf("countSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
