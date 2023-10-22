package offer

import "testing"

func Test_gameOfLife(t *testing.T) {
	type args struct {
		board [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				board: [][]int{{1}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameOfLife(tt.args.board)
		})
	}
}
