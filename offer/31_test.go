package offer

import (
	"fmt"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 3},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			fmt.Println(tt.args.nums)
		})
	}
}
