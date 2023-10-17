package offer

import (
	"fmt"
	"testing"
)

func Test_generate(t *testing.T) {
	type args struct {
		numRows int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				numRows: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generate(tt.args.numRows)
			fmt.Println(got)
		})
	}
}
