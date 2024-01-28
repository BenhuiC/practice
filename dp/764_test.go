package dp

import "testing"

func Test_orderOfLargestPlusSign(t *testing.T) {
	type args struct {
		n     int
		mines [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n:     1,
				mines: [][]int{{0, 0}},
			},
			want: 0,
		},
		{
			name: "2",
			args: args{
				n:     5,
				mines: [][]int{{4, 2}},
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				n:     5,
				mines: [][]int{{3, 0}, {3, 3}},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orderOfLargestPlusSign(tt.args.n, tt.args.mines); got != tt.want {
				t.Errorf("orderOfLargestPlusSign() = %v, want %v", got, tt.want)
			}
		})
	}
}
