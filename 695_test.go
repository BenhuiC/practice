package partice

import "testing"

func Test_maxAreaOfIsland(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name: "1",
			args: args{
				grid: [][]int{
					{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
					{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
					{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}},
			},
			wantRes: 6,
		},
		{
			name: "2",
			args: args{
				grid: [][]int{
					{1, 1, 0, 0, 0},
					{1, 1, 0, 0, 0},
					{0, 0, 0, 1, 1},
					{0, 0, 0, 1, 1}},
			},
			wantRes: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := maxAreaOfIsland(tt.args.grid); gotRes != tt.wantRes {
				t.Errorf("maxAreaOfIsland() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
