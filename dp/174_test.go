package dp

import "testing"

func Test_calculateMinimumHP(t *testing.T) {
	type args struct {
		dungeon [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				dungeon: [][]int{
					{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5},
				},
			},
			want: 7,
		},
		{
			name: "2",
			args: args{
				dungeon: [][]int{
					{0, -5}, {0, 0},
				},
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				dungeon: [][]int{
					{1, -3, 3},
					{0, -2, 0},
					{-3, -3, -3},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMinimumHP(tt.args.dungeon); got != tt.want {
				t.Errorf("calculateMinimumHP() = %v, want %v", got, tt.want)
			}
		})
	}
}
