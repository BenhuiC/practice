package partice

import (
	"fmt"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				board: [][]byte{
					{'O','O','O'},{'O','O','O'},{'O','O','O'},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args.board)
			fmt.Println(tt.args.board)
		})
	}
}
