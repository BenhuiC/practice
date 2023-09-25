package back_trace

import (
	"testing"
)

func Test_permuteUnique(t *testing.T) {
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
				nums: []int{1, 1, 2, 2, 3},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 1, 2},
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{1, 2, 2},
			},
		},
		{
			name: "4",
			args: args{
				nums: []int{0, 0, 1, 1, 2, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permuteUnique2(tt.args.nums)
			t.Log(len(got))
			t.Log(got)
		})
	}
}

/*
 [0 0 3 1 1 2 2] [0 0 3 1 2 1 2] [0 0 3 1 2 2 1] [0 0 3 2 1 1 2] [0 0 3 2 1 2 1] [0 0 3 2 2 1 1] [0 1 0 1 2 2 3] [0 1 0 1 2 3 2] [0 1 0 1 3 2 2] [0 1 0 2 1 2 3] [0 1 0 2 1 3 2] [0 1 0 2 2 1 3] [0 1 0 2 2 3 1] [0 1 0 2 3 1 2] [0 1 0 2 3 2 1] [0 1 0 3 1 2 2] [0 1 0 3 2 1 2] [0 1 0 3 2 2 1] [0 1 1 0 2 2 3] [0 1 1 0 2 3 2] [0 1 1 0 3 2 2] [0 1 1 2 0 2 3] [0 1 1 2 0 3 2] [0 1 1 2 2 0 3] [0 1 1 2 2 3 0] [0 1 1 2 3 0 2] [0 1 1 2 3 2 0] [0 1 1 3 0 2 2] [0 1 1 3 2 0 2] [0 1 1 3 2 2 0] [0 1 2 0 1 2 3] [0 1 2 0 1 3 2] [0 1 2 0 2 1 3] [0 1 2 0 2 3 1] [0 1 2 0 3 1 2] [0 1 2 0 3 2 1] [0 1 2 1 0 2 3] [0 1 2 1 0 3 2] [0 1 2 1 2 0 3] [0 1 2 1 2 3 0] [0 1 2 1 3 0 2] [0 1 2 1 3 2 0] [0 1 2 2 0 1 3] [0 1 2 2 0 3 1] [0 1 2 2 1 0 3] [0 1 2 2 1 3 0] [0 1 2 2 3 0 1] [0 1 2 2 3 1 0] [0 1 2 3 0 1 2] [0 1 2 3 0 2 1] [0 1 2 3 1 0 2] [0 1 2 3 1 2 0] [0 1 2 3 2 0 1] [0 1 2 3 2 1 0] [0 1 3 0 1 2 2] [0 1 3 0 2 1 2] [0 1 3 0 2 2 1] [0 1 3 1 0 2 2] [0 1 3 1 2 0 2] [0 1 3 1 2 2 0] [0 1 3 2 0 1 2] [0 1 3 2 0 2 1] [0 1 3 2 1 0 2] [0 1 3 2 1 2 0] [0 1 3 2 2 0 1] [0 1 3 2 2 1 0] [0 2 0 1 1 2 3] [0 2 0 1 1 3 2] [0 2 0 1 2 1 3] [0 2 0 1 2 3 1] [0 2 0 1 3 1 2] [0 2 0 1 3 2 1] [0 2 0 2 1 1 3] [0 2 0 2 1 3 1] [0 2 0 2 3 1 1] [0 2 0 3 1 1 2] [0 2 0 3 1 2 1] [0 2 0 3 2 1 1] [0 2 1 0 1 2 3] [0 2 1 0 1 3 2] [0 2 1 0 2 1 3] [0 2 1 0 2 3 1] [0 2 1 0 3 1 2] [0 2 1 0 3 2 1] [0 2 1 1 0 2 3] [0 2 1 1 0 3 2] [0 2 1 1 2 0 3] [0 2 1 1 2 3 0] [0 2 1 1 3 0 2] [0 2 1 1 3 2 0] [0 2 1 2 0 1 3] [0 2 1 2 0 3 1] [0 2 1 2 1 0 3] [0 2 1 2 1 3 0] [0 2 1 2 3 0 1] [0 2 1 2 3 1 0] [0 2 1 3 0 1 2] [0 2 1 3 0 2 1] [0 2 1 3 1 0 2] [0 2 1 3 1 2 0] [0 2 1 3 2 0 1] [0 2 1 3 2 1 0] [0 2 2 0 1 1 3] [0 2 2 0 1 3 1] [0 2 2 0 3 1 1] [0 2 2 1 0 1 3] [0 2 2 1 0 3 1] [0 2 2 1 1 0 3] [0 2 2 1 1 3 0] [0 2 2 1 3 0 1] [0 2 2 1 3 1 0] [0 2 2 3 0 1 1] [0 2 2 3 1 0 1] [0 2 2 3 1 1 0] [0 2 3 0 1 1 2] [0 2 3 0 1 2 1] [0 2 3 0 2 1 1] [0 2 3 1 0 1 2] [0 2 3 1 0 2 1] [0 2 3 1 1 0 2] [0 2 3 1 1 2 0] [0 2 3 1 2 0 1] [0 2 3 1 2 1 0] [0 2 3 2 0 1 1] [0 2 3 2 1 0 1] [0 2 3 2 1 1 0] [0 3 0 1 1 2 2] [0 3 0 1 2 1 2] [0 3 0 1 2 2 1] [0 3 0 2 1 1 2] [0 3 0 2 1 2 1] [0 3 0 2 2 1 1] [0 3 1 0 1 2 2] [0 3 1 0 2 1 2] [0 3 1 0 2 2 1] [0 3 1 1 0 2 2] [0 3 1 1 2 0 2] [0 3 1 1 2 2 0] [0 3 1 2 0 1 2] [0 3 1 2 0 2 1] [0 3 1 2 1 0 2] [0 3 1 2 1 2 0] [0 3 1 2 2 0 1] [0 3 1 2 2 1 0] [0 3 2 0 1 1 2] [0 3 2 0 1 2 1] [0 3 2 0 2 1 1] [0 3 2 1 0 1 2] [0 3 2 1 0 2 1] [0 3 2 1 1 0 2] [0 3 2 1 1 2 0] [0 3 2 1 2 0 1] [0 3 2 1 2 1 0] [0 3 2 2 0 1 1] [0 3 2 2 1 0 1] [0 3 2 2 1 1 0] [1 0 0 1 2 2 3] [1 0 0 1 2 3 2] [1 0 0 1 3 2 2] [1 0 0 2 1 2 3] [1 0 0 2 1 3 2] [1 0 0 2 2 1 3] [1 0 0 2 2 3 1] [1 0 0 2 3 1 2] [1 0 0 2 3 2 1] [1 0 0 3 1 2 2] [1 0 0 3 2 1 2] [1 0 0 3 2 2 1] [1 0 1 0 2 2 3] [1 0 1 0 2 3 2] [1 0 1 0 3 2 2] [1 0 1 2 0 2 3] [1 0 1 2 0 3 2] [1 0 1 2 2 0 3] [1 0 1 2 2 3 0] [1 0 1 2 3 0 2] [1 0 1 2 3 2 0] [1 0 1 3 0 2 2] [1 0 1 3 2 0 2] [1 0 1 3 2 2 0] [1 0 2 0 1 2 3] [1 0 2 0 1 3 2] [1 0 2 0 2 1 3] [1 0 2 0 2 3 1] [1 0 2 0 3 1 2] [1 0 2 0 3 2 1] [1 0 2 1 0 2 3] [1 0 2 1 0 3 2] [1 0 2 1 2 0 3] [1 0 2 1 2 3 0] [1 0 2 1 3 0 2] [1 0 2 1 3 2 0] [1 0 2 2 0 1 3] [1 0 2 2 0 3 1] [1 0 2 2 1 0 3] [1 0 2 2 1 3 0] [1 0 2 2 3 0 1] [1 0 2 2 3 1 0] [1 0 2 3 0 1 2] [1 0 2 3 0 2 1] [1 0 2 3 1 0 2] [1 0 2 3 1 2 0] [1 0 2 3 2 0 1] [1 0 2 3 2 1 0] [1 0 3 0 1 2 2] [1 0 3 0 2 1 2] [1 0 3 0 2 2 1] [1 0 3 1 0 2 2] [1 0 3 1 2 0 2] [1 0 3 1 2 2 0] [1 0 3 2 0 1 2] [1 0 3 2 0 2 1] [1 0 3 2 1 0 2] [1 0 3 2 1 2 0] [1 0 3 2 2 0 1] [1 0 3 2 2 1 0] [1 1 0 0 2 2 3] [1 1 0 0 2 3 2] [1 1 0 0 3 2 2] [1 1 0 2 0 2 3] [1 1 0 2 0 3 2] [1 1 0 2 2 0 3] [1 1 0 2 2 3 0] [1 1 0 2 3 0 2] [1 1 0 2 3 2 0] [1 1 0 3 0 2 2] [1 1 0 3 2 0 2] [1 1 0 3 2 2 0] [1 1 2 0 0 2 3] [1 1 2 0 0 3 2] [1 1 2 0 2 0 3] [1 1 2 0 2 3 0] [1 1 2 0 3 0 2] [1 1 2 0 3 2 0] [1 1 2 2 0 0 3] [1 1 2 2 0 3 0] [1 1 2 2 3 0 0] [1 1 2 3 0 0 2] [1 1 2 3 0 2 0] [1 1 2 3 2 0 0] [1 1 3 0 0 2 2] [1 1 3 0 2 0 2] [1 1 3 0 2 2 0] [1 1 3 2 0 0 2] [1 1 3 2 0 2 0] [1 1 3 2 2 0 0] [1 2 0 0 1 2 3] [1 2 0 0 1 3 2] [1 2 0 0 2 1 3] [1 2 0 0 2 3 1] [1 2 0 0 3 1 2] [1 2 0 0 3 2 1] [1 2 0 1 0 2 3] [1 2 0 1 0 3 2] [1 2 0 1 2 0 3] [1 2 0 1 2 3 0] [1 2 0 1 3 0 2] [1 2 0 1 3 2 0] [1 2 0 2 0 1 3] [1 2 0 2 0 3 1] [1 2 0 2 1 0 3] [1 2 0 2 1 3 0] [1 2 0 2 3 0 1] [1 2 0 2 3 1 0] [1 2 0 3 0 1 2] [1 2 0 3 0 2 1] [1 2 0 3 1 0 2] [1 2 0 3 1 2 0] [1 2 0 3 2 0 1] [1 2 0 3 2 1 0] [1 2 1 0 0 2 3] [1 2 1 0 0 3 2] [1 2 1 0 2 0 3] [1 2 1 0 2 3 0] [1 2 1 0 3 0 2] [1 2 1 0 3 2 0] [1 2 1 2 0 0 3] [1 2 1 2 0 3 0] [1 2 1 2 3 0 0] [1 2 1 3 0 0 2] [1 2 1 3 0 2 0] [1 2 1 3 2 0 0] [1 2 2 0 0 1 3] [1 2 2 0 0 3 1] [1 2 2 0 1 0 3] [1 2 2 0 1 3 0] [1 2 2 0 3 0 1] [1 2 2 0 3 1 0] [1 2 2 1 0 0 3] [1 2 2 1 0 3 0] [1 2 2 1 3 0 0] [1 2 2 3 0 0 1] [1 2 2 3 0 1 0] [1 2 2 3 1 0 0] [1 2 3 0 0 1 2] [1 2 3 0 0 2 1] [1 2 3 0 1 0 2] [1 2 3 0 1 2 0] [1 2 3 0 2 0 1] [1 2 3 0 2 1 0] [1 2 3 1 0 0 2] [1 2 3 1 0 2 0] [1 2 3 1 2 0 0] [1 2 3 2 0 0 1] [1 2 3 2 0 1 0] [1 2 3 2 1 0 0] [1 3 0 0 1 2 2] [1 3 0 0 2 1 2] [1 3 0 0 2 2 1] [1 3 0 1 0 2 2] [1 3 0 1 2 0 2] [1 3 0 1 2 2 0] [1 3 0 2 0 1 2] [1 3 0 2 0 2 1] [1 3 0 2 1 0 2] [1 3 0 2 1 2 0] [1 3 0 2 2 0 1] [1 3 0 2 2 1 0] [1 3 1 0 0 2 2] [1 3 1 0 2 0 2] [1 3 1 0 2 2 0] [1 3 1 2 0 0 2] [1 3 1 2 0 2 0] [1 3 1 2 2 0 0] [1 3 2 0 0 1 2] [1 3 2 0 0 2 1] [1 3 2 0 1 0 2] [1 3 2 0 1 2 0] [1 3 2 0 2 0 1] [1 3 2 0 2 1 0] [1 3 2 1 0 0 2] [1 3 2 1 0 2 0] [1 3 2 1 2 0 0] [1 3 2 2 0 0 1] [1 3 2 2 0 1 0] [1 3 2 2 1 0 0] [2 0 0 1 1 2 3] [2 0 0 1 1 3 2] [2 0 0 1 2 1 3] [2 0 0 1 2 3 1] [2 0 0 1 3 1 2] [2 0 0 1 3 2 1] [2 0 0 2 1 1 3] [2 0 0 2 1 3 1] [2 0 0 2 3 1 1] [2 0 0 3 1 1 2] [2 0 0 3 1 2 1] [2 0 0 3 2 1 1] [2 0 1 0 1 2 3] [2 0 1 0 1 3 2] [2 0 1 0 2 1 3] [2 0 1 0 2 3 1] [2 0 1 0 3 1 2] [2 0 1 0 3 2 1] [2 0 1 1 0 2 3] [2 0 1 1 0 3 2] [2 0 1 1 2 0 3] [2 0 1 1 2 3 0] [2 0 1 1 3 0 2] [2 0 1 1 3 2 0] [2 0 1 2 0 1 3] [2 0 1 2 0 3 1] [2 0 1 2 1 0 3] [2 0 1 2 1 3 0] [2 0 1 2 3 0 1] [2 0 1 2 3 1 0] [2 0 1 3 0 1 2] [2 0 1 3 0 2 1] [2 0 1 3 1 0 2] [2 0 1 3 1 2 0] [2 0 1 3 2 0 1] [2 0 1 3 2 1 0] [2 0 2 0 1 1 3] [2 0 2 0 1 3 1] [2 0 2 0 3 1 1] [2 0 2 1 0 1 3] [2 0 2 1 0 3 1] [2 0 2 1 1 0 3] [2 0 2 1 1 3 0] [2 0 2 1 3 0 1] [2 0 2 1 3 1 0] [2 0 2 3 0 1 1] [2 0 2 3 1 0 1] [2 0 2 3 1 1 0] [2 0 3 0 1 1 2] [2 0 3 0 1 2 1] [2 0 3 0 2 1 1] [2 0 3 1 0 1 2] [2 0 3 1 0 2 1] [2 0 3 1 1 0 2] [2 0 3 1 1 2 0] [2 0 3 1 2 0 1] [2 0 3 1 2 1 0] [2 0 3 2 0 1 1] [2 0 3 2 1 0 1] [2 0 3 2 1 1 0] [2 1 0 0 1 2 3] [2 1 0 0 1 3 2] [2 1 0 0 2 1 3] [2 1 0 0 2 3 1] [2 1 0 0 3 1 2] [2 1 0 0 3 2 1] [2 1 0 1 0 2 3] [2 1 0 1 0 3 2] [2 1 0 1 2 0 3] [2 1 0 1 2 3 0] [2 1 0 1 3 0 2] [2 1 0 1 3 2 0] [2 1 0 2 0 1 3] [2 1 0 2 0 3 1] [2 1 0 2 1 0 3] [2 1 0 2 1 3 0] [2 1 0 2 3 0 1] [2 1 0 2 3 1 0] [2 1 0 3 0 1 2] [2 1 0 3 0 2 1] [2 1 0 3 1 0 2] [2 1 0 3 1 2 0] [2 1 0 3 2 0 1] [2 1 0 3 2 1 0] [2 1 1 0 0 2 3] [2 1 1 0 0 3 2] [2 1 1 0 2 0 3] [2 1 1 0 2 3 0] [2 1 1 0 3 0 2] [2 1 1 0 3 2 0] [2 1 1 2 0 0 3] [2 1 1 2 0 3 0] [2 1 1 2 3 0 0] [2 1 1 3 0 0 2] [2 1 1 3 0 2 0] [2 1 1 3 2 0 0] [2 1 2 0 0 1 3] [2 1 2 0 0 3 1] [2 1 2 0 1 0 3] [2 1 2 0 1 3 0] [2 1 2 0 3 0 1] [2 1 2 0 3 1 0] [2 1 2 1 0 0 3] [2 1 2 1 0 3 0] [2 1 2 1 3 0 0] [2 1 2 3 0 0 1] [2 1 2 3 0 1 0] [2 1 2 3 1 0 0] [2 1 3 0 0 1 2] [2 1 3 0 0 2 1] [2 1 3 0 1 0 2] [2 1 3 0 1 2 0] [2 1 3 0 2 0 1] [2 1 3 0 2 1 0] [2 1 3 1 0 0 2] [2 1 3 1 0 2 0] [2 1 3 1 2 0 0] [2 1 3 2 0 0 1] [2 1 3 2 0 1 0] [2 1 3 2 1 0 0] [2 2 0 0 1 1 3] [2 2 0 0 1 3 1] [2 2 0 0 3 1 1] [2 2 0 1 0 1 3] [2 2 0 1 0 3 1] [2 2 0 1 1 0 3] [2 2 0 1 1 3 0] [2 2 0 1 3 0 1] [2 2 0 1 3 1 0] [2 2 0 3 0 1 1] [2 2 0 3 1 0 1] [2 2 0 3 1 1 0] [2 2 1 0 0 1 3] [2 2 1 0 0 3 1] [2 2 1 0 1 0 3] [2 2 1 0 1 3 0] [2 2 1 0 3 0 1] [2 2 1 0 3 1 0] [2 2 1 1 0 0 3] [2 2 1 1 0 3 0] [2 2 1 1 3 0 0] [2 2 1 3 0 0 1] [2 2 1 3 0 1 0] [2 2 1 3 1 0 0] [2 2 3 0 0 1 1] [2 2 3 0 1 0 1] [2 2 3 0 1 1 0] [2 2 3 1 0 0 1] [2 2 3 1 0 1 0] [2 2 3 1 1 0 0] [2 3 0 0 1 1 2] [2 3 0 0 1 2 1] [2 3 0 0 2 1 1] [2 3 0 1 0 1 2] [2 3 0 1 0 2 1] [2 3 0 1 1 0 2] [2 3 0 1 1 2 0] [2 3 0 1 2 0 1] [2 3 0 1 2 1 0] [2 3 0 2 0 1 1] [2 3 0 2 1 0 1] [2 3 0 2 1 1 0] [2 3 1 0 0 1 2] [2 3 1 0 0 2 1] [2 3 1 0 1 0 2] [2 3 1 0 1 2 0] [2 3 1 0 2 0 1] [2 3 1 0 2 1 0] [2 3 1 1 0 0 2] [2 3 1 1 0 2 0] [2 3 1 1 2 0 0] [2 3 1 2 0 0 1] [2 3 1 2 0 1 0] [2 3 1 2 1 0 0] [2 3 2 0 0 1 1] [2 3 2 0 1 0 1] [2 3 2 0 1 1 0] [2 3 2 1 0 0 1] [2 3 2 1 0 1 0] [2 3 2 1 1 0 0] [3 0 0 1 1 2 2] [3 0 0 1 2 1 2] [3 0 0 1 2 2 1] [3 0 0 2 1 1 2] [3 0 0 2 1 2 1] [3 0 0 2 2 1 1] [3 0 1 0 1 2 2] [3 0 1 0 2 1 2] [3 0 1 0 2 2 1] [3 0 1 1 0 2 2] [3 0 1 1 2 0 2] [3 0 1 1 2 2 0] [3 0 1 2 0 1 2] [3 0 1 2 0 2 1] [3 0 1 2 1 0 2] [3 0 1 2 1 2 0] [3 0 1 2 2 0 1] [3 0 1 2 2 1 0] [3 0 2 0 1 1 2] [3 0 2 0 1 2 1] [3 0 2 0 2 1 1] [3 0 2 1 0 1 2] [3 0 2 1 0 2 1] [3 0 2 1 1 0 2] [3 0 2 1 1 2 0] [3 0 2 1 2 0 1] [3 0 2 1 2 1 0] [3 0 2 2 0 1 1] [3 0 2 2 1 0 1] [3 0 2 2 1 1 0] [3 1 0 0 1 2 2] [3 1 0 0 2 1 2] [3 1 0 0 2 2 1] [3 1 0 1 0 2 2] [3 1 0 1 2 0 2] [3 1 0 1 2 2 0] [3 1 0 2 0 1 2] [3 1 0 2 0 2 1] [3 1 0 2 1 0 2] [3 1 0 2 1 2 0] [3 1 0 2 2 0 1] [3 1 0 2 2 1 0] [3 1 1 0 0 2 2] [3 1 1 0 2 0 2] [3 1 1 0 2 2 0] [3 1 1 2 0 0 2] [3 1 1 2 0 2 0] [3 1 1 2 2 0 0] [3 1 2 0 0 1 2] [3 1 2 0 0 2 1] [3 1 2 0 1 0 2] [3 1 2 0 1 2 0] [3 1 2 0 2 0 1] [3 1 2 0 2 1 0] [3 1 2 1 0 0 2] [3 1 2 1 0 2 0] [3 1 2 1 2 0 0] [3 1 2 2 0 0 1] [3 1 2 2 0 1 0] [3 1 2 2 1 0 0] [3 2 0 0 1 1 2] [3 2 0 0 1 2 1] [3 2 0 0 2 1 1] [3 2 0 1 0 1 2] [3 2 0 1 0 2 1] [3 2 0 1 1 0 2] [3 2 0 1 1 2 0] [3 2 0 1 2 0 1] [3 2 0 1 2 1 0] [3 2 0 2 0 1 1] [3 2 0 2 1 0 1] [3 2 0 2 1 1 0] [3 2 1 0 0 1 2] [3 2 1 0 0 2 1] [3 2 1 0 1 0 2] [3 2 1 0 1 2 0] [3 2 1 0 2 0 1] [3 2 1 0 2 1 0] [3 2 1 1 0 0 2] [3 2 1 1 0 2 0] [3 2 1 1 2 0 0] [3 2 1 2 0 0 1] [3 2 1 2 0 1 0] [3 2 1 2 1 0 0] [3 2 2 0 0 1 1] [3 2 2 0 1 0 1] [3 2 2 0 1 1 0] [3 2 2 1 0 0 1] [3 2 2 1 0 1 0] [3 2 2 1 1 0 0]]
 [0 0 3 1 2 2 1] [0 0 3 1 2 1 2] [0 0 3 1 1 2 2] [0 0 3 2 1 2 1] [0 0 3 2 1 1 2] [0 0 3 2 2 1 1] [0 1 0 1 2 2 3] [0 1 0 1 2 3 2] [0 1 0 1 3 2 2] [0 1 0 2 1 2 3] [0 1 0 2 1 3 2] [0 1 0 2 2 1 3] [0 1 0 2 2 3 1] [0 1 0 2 3 2 1] [0 1 0 2 3 1 2] [0 1 0 3 2 2 1] [0 1 0 3 2 1 2] [0 1 0 3 1 2 2] [0 1 1 0 2 2 3] [0 1 1 0 2 3 2] [0 1 1 0 3 2 2] [0 1 1 2 0 2 3] [0 1 1 2 0 3 2] [0 1 1 2 2 0 3] [0 1 1 2 2 3 0] [0 1 1 2 3 2 0] [0 1 1 2 3 0 2] [0 1 1 3 2 2 0] [0 1 1 3 2 0 2] [0 1 1 3 0 2 2] [0 1 2 1 0 2 3] [0 1 2 1 0 3 2] [0 1 2 1 2 0 3] [0 1 2 1 2 3 0] [0 1 2 1 3 2 0] [0 1 2 1 3 0 2] [0 1 2 0 1 2 3] [0 1 2 0 1 3 2] [0 1 2 0 2 1 3] [0 1 2 0 2 3 1] [0 1 2 0 3 2 1] [0 1 2 0 3 1 2] [0 1 2 2 0 1 3] [0 1 2 2 0 3 1] [0 1 2 2 1 0 3] [0 1 2 2 1 3 0] [0 1 2 2 3 1 0] [0 1 2 2 3 0 1] [0 1 2 3 0 2 1] [0 1 2 3 0 1 2] [0 1 2 3 2 0 1] [0 1 2 3 2 1 0] [0 1 2 3 1 2 0] [0 1 2 3 1 0 2] [0 1 3 1 2 2 0] [0 1 3 1 2 0 2] [0 1 3 1 0 2 2] [0 1 3 2 1 2 0] [0 1 3 2 1 0 2] [0 1 3 2 2 1 0] [0 1 3 2 2 0 1] [0 1 3 2 0 2 1] [0 1 3 2 0 1 2] [0 1 3 0 2 2 1] [0 1 3 0 2 1 2] [0 1 3 0 1 2 2] [0 2 1 1 0 2 3] [0 2 1 1 0 3 2] [0 2 1 1 2 0 3] [0 2 1 1 2 3 0] [0 2 1 1 3 2 0] [0 2 1 1 3 0 2] [0 2 1 0 1 2 3] [0 2 1 0 1 3 2] [0 2 1 0 2 1 3] [0 2 1 0 2 3 1] [0 2 1 0 3 2 1] [0 2 1 0 3 1 2] [0 2 1 2 0 1 3] [0 2 1 2 0 3 1] [0 2 1 2 1 0 3] [0 2 1 2 1 3 0] [0 2 1 2 3 1 0] [0 2 1 2 3 0 1] [0 2 1 3 0 2 1] [0 2 1 3 0 1 2] [0 2 1 3 2 0 1] [0 2 1 3 2 1 0] [0 2 1 3 1 2 0] [0 2 1 3 1 0 2] [0 2 0 1 1 2 3] [0 2 0 1 1 3 2] [0 2 0 1 2 1 3] [0 2 0 1 2 3 1] [0 2 0 1 3 2 1] [0 2 0 1 3 1 2] [0 2 0 2 1 1 3] [0 2 0 2 1 3 1] [0 2 0 2 3 1 1] [0 2 0 3 1 2 1] [0 2 0 3 1 1 2] [0 2 0 3 2 1 1] [0 2 2 1 0 1 3] [0 2 2 1 0 3 1] [0 2 2 1 1 0 3] [0 2 2 1 1 3 0] [0 2 2 1 3 1 0] [0 2 2 1 3 0 1] [0 2 2 0 1 1 3] [0 2 2 0 1 3 1] [0 2 2 0 3 1 1] [0 2 2 3 0 1 1] [0 2 2 3 1 0 1] [0 2 2 3 1 1 0] [0 2 3 1 0 2 1] [0 2 3 1 0 1 2] [0 2 3 1 2 0 1] [0 2 3 1 2 1 0] [0 2 3 1 1 2 0] [0 2 3 1 1 0 2] [0 2 3 0 1 2 1] [0 2 3 0 1 1 2] [0 2 3 0 2 1 1] [0 2 3 2 0 1 1] [0 2 3 2 1 0 1] [0 2 3 2 1 1 0] [0 3 1 1 2 2 0] [0 3 1 1 2 0 2] [0 3 1 1 0 2 2] [0 3 1 2 1 2 0] [0 3 1 2 1 0 2] [0 3 1 2 2 1 0] [0 3 1 2 2 0 1] [0 3 1 2 0 2 1] [0 3 1 2 0 1 2] [0 3 1 0 2 2 1] [0 3 1 0 2 1 2] [0 3 1 0 1 2 2] [0 3 2 1 1 2 0] [0 3 2 1 1 0 2] [0 3 2 1 2 1 0] [0 3 2 1 2 0 1] [0 3 2 1 0 2 1] [0 3 2 1 0 1 2] [0 3 2 2 1 1 0] [0 3 2 2 1 0 1] [0 3 2 2 0 1 1] [0 3 2 0 1 2 1] [0 3 2 0 1 1 2] [0 3 2 0 2 1 1] [0 3 0 1 2 2 1] [0 3 0 1 2 1 2] [0 3 0 1 1 2 2] [0 3 0 2 1 2 1] [0 3 0 2 1 1 2] [0 3 0 2 2 1 1] [1 0 0 1 2 2 3] [1 0 0 1 2 3 2] [1 0 0 1 3 2 2] [1 0 0 2 1 2 3] [1 0 0 2 1 3 2] [1 0 0 2 2 1 3] [1 0 0 2 2 3 1] [1 0 0 2 3 2 1] [1 0 0 2 3 1 2] [1 0 0 3 2 2 1] [1 0 0 3 2 1 2] [1 0 0 3 1 2 2] [1 0 1 0 2 2 3] [1 0 1 0 2 3 2] [1 0 1 0 3 2 2] [1 0 1 2 0 2 3] [1 0 1 2 0 3 2] [1 0 1 2 2 0 3] [1 0 1 2 2 3 0] [1 0 1 2 3 2 0] [1 0 1 2 3 0 2] [1 0 1 3 2 2 0] [1 0 1 3 2 0 2] [1 0 1 3 0 2 2] [1 0 2 1 0 2 3] [1 0 2 1 0 3 2] [1 0 2 1 2 0 3] [1 0 2 1 2 3 0] [1 0 2 1 3 2 0] [1 0 2 1 3 0 2] [1 0 2 0 1 2 3] [1 0 2 0 1 3 2] [1 0 2 0 2 1 3] [1 0 2 0 2 3 1] [1 0 2 0 3 2 1] [1 0 2 0 3 1 2] [1 0 2 2 0 1 3] [1 0 2 2 0 3 1] [1 0 2 2 1 0 3] [1 0 2 2 1 3 0] [1 0 2 2 3 1 0] [1 0 2 2 3 0 1] [1 0 2 3 0 2 1] [1 0 2 3 0 1 2] [1 0 2 3 2 0 1] [1 0 2 3 2 1 0] [1 0 2 3 1 2 0] [1 0 2 3 1 0 2] [1 0 3 1 2 2 0] [1 0 3 1 2 0 2] [1 0 3 1 0 2 2] [1 0 3 2 1 2 0] [1 0 3 2 1 0 2] [1 0 3 2 2 1 0] [1 0 3 2 2 0 1] [1 0 3 2 0 2 1] [1 0 3 2 0 1 2] [1 0 3 0 2 2 1] [1 0 3 0 2 1 2] [1 0 3 0 1 2 2] [1 1 0 0 2 2 3] [1 1 0 0 2 3 2] [1 1 0 0 3 2 2] [1 1 0 2 0 2 3] [1 1 0 2 0 3 2] [1 1 0 2 2 0 3] [1 1 0 2 2 3 0] [1 1 0 2 3 2 0] [1 1 0 2 3 0 2] [1 1 0 3 2 2 0] [1 1 0 3 2 0 2] [1 1 0 3 0 2 2] [1 1 2 0 0 2 3] [1 1 2 0 0 3 2] [1 1 2 0 2 0 3] [1 1 2 0 2 3 0] [1 1 2 0 3 2 0] [1 1 2 0 3 0 2] [1 1 2 2 0 0 3] [1 1 2 2 0 3 0] [1 1 2 2 3 0 0] [1 1 2 3 0 2 0] [1 1 2 3 0 0 2] [1 1 2 3 2 0 0] [1 1 3 0 2 2 0] [1 1 3 0 2 0 2] [1 1 3 0 0 2 2] [1 1 3 2 0 2 0] [1 1 3 2 0 0 2] [1 1 3 2 2 0 0] [1 2 0 1 0 2 3] [1 2 0 1 0 3 2] [1 2 0 1 2 0 3] [1 2 0 1 2 3 0] [1 2 0 1 3 2 0] [1 2 0 1 3 0 2] [1 2 0 0 1 2 3] [1 2 0 0 1 3 2] [1 2 0 0 2 1 3] [1 2 0 0 2 3 1] [1 2 0 0 3 2 1] [1 2 0 0 3 1 2] [1 2 0 2 0 1 3] [1 2 0 2 0 3 1] [1 2 0 2 1 0 3] [1 2 0 2 1 3 0] [1 2 0 2 3 1 0] [1 2 0 2 3 0 1] [1 2 0 3 0 2 1] [1 2 0 3 0 1 2] [1 2 0 3 2 0 1] [1 2 0 3 2 1 0] [1 2 0 3 1 2 0] [1 2 0 3 1 0 2] [1 2 1 0 0 2 3] [1 2 1 0 0 3 2] [1 2 1 0 2 0 3] [1 2 1 0 2 3 0] [1 2 1 0 3 2 0] [1 2 1 0 3 0 2] [1 2 1 2 0 0 3] [1 2 1 2 0 3 0] [1 2 1 2 3 0 0] [1 2 1 3 0 2 0] [1 2 1 3 0 0 2] [1 2 1 3 2 0 0] [1 2 2 1 0 0 3] [1 2 2 1 0 3 0] [1 2 2 1 3 0 0] [1 2 2 0 1 0 3] [1 2 2 0 1 3 0] [1 2 2 0 0 1 3] [1 2 2 0 0 3 1] [1 2 2 0 3 0 1] [1 2 2 0 3 1 0] [1 2 2 3 0 0 1] [1 2 2 3 0 1 0] [1 2 2 3 1 0 0] [1 2 3 1 0 2 0] [1 2 3 1 0 0 2] [1 2 3 1 2 0 0] [1 2 3 0 1 2 0] [1 2 3 0 1 0 2] [1 2 3 0 2 1 0] [1 2 3 0 2 0 1] [1 2 3 0 0 2 1] [1 2 3 0 0 1 2] [1 2 3 2 0 1 0] [1 2 3 2 0 0 1] [1 2 3 2 1 0 0] [1 2 3 0 0 2 1] [1 2 3 0 0 1 2] [1 2 3 0 2 0 1] [1 2 3 0 2 1 0] [1 2 3 0 1 2 0] [1 2 3 0 1 0 2] [1 3 0 1 2 2 0] [1 3 0 1 2 0 2] [1 3 0 1 0 2 2] [1 3 0 2 1 2 0] [1 3 0 2 1 0 2] [1 3 0 2 2 1 0] [1 3 0 2 2 0 1] [1 3 0 2 0 2 1] [1 3 0 2 0 1 2] [1 3 0 0 2 2 1] [1 3 0 0 2 1 2] [1 3 0 0 1 2 2] [1 3 1 0 2 2 0] [1 3 1 0 2 0 2] [1 3 1 0 0 2 2] [1 3 1 2 0 2 0] [1 3 1 2 0 0 2] [1 3 1 2 2 0 0] [1 3 2 1 0 2 0] [1 3 2 1 0 0 2] [1 3 2 1 2 0 0] [1 3 2 0 1 2 0] [1 3 2 0 1 0 2] [1 3 2 0 2 1 0] [1 3 2 0 2 0 1] [1 3 2 0 0 2 1] [1 3 2 0 0 1 2] [1 3 2 2 0 1 0] [1 3 2 2 0 0 1] [1 3 2 2 1 0 0] [1 3 2 0 0 2 1] [1 3 2 0 0 1 2] [1 3 2 0 2 0 1] [1 3 2 0 2 1 0] [1 3 2 0 1 2 0] [1 3 2 0 1 0 2] [2 0 1 1 0 2 3] [2 0 1 1 0 3 2] [2 0 1 1 2 0 3] [2 0 1 1 2 3 0] [2 0 1 1 3 2 0] [2 0 1 1 3 0 2] [2 0 1 0 1 2 3] [2 0 1 0 1 3 2] [2 0 1 0 2 1 3] [2 0 1 0 2 3 1] [2 0 1 0 3 2 1] [2 0 1 0 3 1 2] [2 0 1 2 0 1 3] [2 0 1 2 0 3 1] [2 0 1 2 1 0 3] [2 0 1 2 1 3 0] [2 0 1 2 3 1 0] [2 0 1 2 3 0 1] [2 0 1 3 0 2 1] [2 0 1 3 0 1 2] [2 0 1 3 2 0 1] [2 0 1 3 2 1 0] [2 0 1 3 1 2 0] [2 0 1 3 1 0 2] [2 0 0 1 1 2 3] [2 0 0 1 1 3 2] [2 0 0 1 2 1 3] [2 0 0 1 2 3 1] [2 0 0 1 3 2 1] [2 0 0 1 3 1 2] [2 0 0 2 1 1 3] [2 0 0 2 1 3 1] [2 0 0 2 3 1 1] [2 0 0 3 1 2 1] [2 0 0 3 1 1 2] [2 0 0 3 2 1 1] [2 0 2 1 0 1 3] [2 0 2 1 0 3 1] [2 0 2 1 1 0 3] [2 0 2 1 1 3 0] [2 0 2 1 3 1 0] [2 0 2 1 3 0 1] [2 0 2 0 1 1 3] [2 0 2 0 1 3 1] [2 0 2 0 3 1 1] [2 0 2 3 0 1 1] [2 0 2 3 1 0 1] [2 0 2 3 1 1 0] [2 0 3 1 0 2 1] [2 0 3 1 0 1 2] [2 0 3 1 2 0 1] [2 0 3 1 2 1 0] [2 0 3 1 1 2 0] [2 0 3 1 1 0 2] [2 0 3 0 1 2 1] [2 0 3 0 1 1 2] [2 0 3 0 2 1 1] [2 0 3 2 0 1 1] [2 0 3 2 1 0 1] [2 0 3 2 1 1 0] [2 1 0 1 0 2 3] [2 1 0 1 0 3 2] [2 1 0 1 2 0 3] [2 1 0 1 2 3 0] [2 1 0 1 3 2 0] [2 1 0 1 3 0 2] [2 1 0 0 1 2 3] [2 1 0 0 1 3 2] [2 1 0 0 2 1 3] [2 1 0 0 2 3 1] [2 1 0 0 3 2 1] [2 1 0 0 3 1 2] [2 1 0 2 0 1 3] [2 1 0 2 0 3 1] [2 1 0 2 1 0 3] [2 1 0 2 1 3 0] [2 1 0 2 3 1 0] [2 1 0 2 3 0 1] [2 1 0 3 0 2 1] [2 1 0 3 0 1 2] [2 1 0 3 2 0 1] [2 1 0 3 2 1 0] [2 1 0 3 1 2 0] [2 1 0 3 1 0 2] [2 1 1 0 0 2 3] [2 1 1 0 0 3 2] [2 1 1 0 2 0 3] [2 1 1 0 2 3 0] [2 1 1 0 3 2 0] [2 1 1 0 3 0 2] [2 1 1 2 0 0 3] [2 1 1 2 0 3 0] [2 1 1 2 3 0 0] [2 1 1 3 0 2 0] [2 1 1 3 0 0 2] [2 1 1 3 2 0 0] [2 1 2 1 0 0 3] [2 1 2 1 0 3 0] [2 1 2 1 3 0 0] [2 1 2 0 1 0 3] [2 1 2 0 1 3 0] [2 1 2 0 0 1 3] [2 1 2 0 0 3 1] [2 1 2 0 3 0 1] [2 1 2 0 3 1 0] [2 1 2 3 0 0 1] [2 1 2 3 0 1 0] [2 1 2 3 1 0 0] [2 1 3 1 0 2 0] [2 1 3 1 0 0 2] [2 1 3 1 2 0 0] [2 1 3 0 1 2 0] [2 1 3 0 1 0 2] [2 1 3 0 2 1 0] [2 1 3 0 2 0 1] [2 1 3 0 0 2 1] [2 1 3 0 0 1 2] [2 1 3 2 0 1 0] [2 1 3 2 0 0 1] [2 1 3 2 1 0 0] [2 1 3 0 0 2 1] [2 1 3 0 0 1 2] [2 1 3 0 2 0 1] [2 1 3 0 2 1 0] [2 1 3 0 1 2 0] [2 1 3 0 1 0 2] [2 2 1 1 0 0 3] [2 2 1 1 0 3 0] [2 2 1 1 3 0 0] [2 2 1 0 1 0 3] [2 2 1 0 1 3 0] [2 2 1 0 0 1 3] [2 2 1 0 0 3 1] [2 2 1 0 3 0 1] [2 2 1 0 3 1 0] [2 2 1 3 0 0 1] [2 2 1 3 0 1 0] [2 2 1 3 1 0 0] [2 2 0 1 1 0 3] [2 2 0 1 1 3 0] [2 2 0 1 0 1 3] [2 2 0 1 0 3 1] [2 2 0 1 3 0 1] [2 2 0 1 3 1 0] [2 2 0 0 1 1 3] [2 2 0 0 1 3 1] [2 2 0 0 3 1 1] [2 2 0 3 1 0 1] [2 2 0 3 1 1 0] [2 2 0 3 0 1 1] [2 2 3 1 0 0 1] [2 2 3 1 0 1 0] [2 2 3 1 1 0 0] [2 2 3 0 1 0 1] [2 2 3 0 1 1 0] [2 2 3 0 0 1 1] [2 3 1 1 0 2 0] [2 3 1 1 0 0 2] [2 3 1 1 2 0 0] [2 3 1 0 1 2 0] [2 3 1 0 1 0 2] [2 3 1 0 2 1 0] [2 3 1 0 2 0 1] [2 3 1 0 0 2 1] [2 3 1 0 0 1 2] [2 3 1 2 0 1 0] [2 3 1 2 0 0 1] [2 3 1 2 1 0 0] [2 3 1 0 0 2 1] [2 3 1 0 0 1 2] [2 3 1 0 2 0 1] [2 3 1 0 2 1 0] [2 3 1 0 1 2 0] [2 3 1 0 1 0 2] [2 3 0 1 1 2 0] [2 3 0 1 1 0 2] [2 3 0 1 2 1 0] [2 3 0 1 2 0 1] [2 3 0 1 0 2 1] [2 3 0 1 0 1 2] [2 3 0 2 1 1 0] [2 3 0 2 1 0 1] [2 3 0 2 0 1 1] [2 3 0 0 1 2 1] [2 3 0 0 1 1 2] [2 3 0 0 2 1 1] [2 3 2 1 0 1 0] [2 3 2 1 0 0 1] [2 3 2 1 1 0 0] [2 3 2 0 1 1 0] [2 3 2 0 1 0 1] [2 3 2 0 0 1 1] [2 3 2 0 0 1 1] [2 3 2 0 1 0 1] [2 3 2 0 1 1 0] [2 3 0 1 0 2 1] [2 3 0 1 0 1 2] [2 3 0 1 2 0 1] [2 3 0 1 2 1 0] [2 3 0 1 1 2 0] [2 3 0 1 1 0 2] [2 3 0 0 1 2 1] [2 3 0 0 1 1 2] [2 3 0 0 2 1 1] [2 3 0 2 0 1 1] [2 3 0 2 1 0 1] [2 3 0 2 1 1 0] [3 0 1 1 2 2 0] [3 0 1 1 2 0 2] [3 0 1 1 0 2 2] [3 0 1 2 1 2 0] [3 0 1 2 1 0 2] [3 0 1 2 2 1 0] [3 0 1 2 2 0 1] [3 0 1 2 0 2 1] [3 0 1 2 0 1 2] [3 0 1 0 2 2 1] [3 0 1 0 2 1 2] [3 0 1 0 1 2 2] [3 0 2 1 1 2 0] [3 0 2 1 1 0 2] [3 0 2 1 2 1 0] [3 0 2 1 2 0 1] [3 0 2 1 0 2 1] [3 0 2 1 0 1 2] [3 0 2 2 1 1 0] [3 0 2 2 1 0 1] [3 0 2 2 0 1 1] [3 0 2 0 1 2 1] [3 0 2 0 1 1 2] [3 0 2 0 2 1 1] [3 0 0 1 2 2 1] [3 0 0 1 2 1 2] [3 0 0 1 1 2 2] [3 0 0 2 1 2 1] [3 0 0 2 1 1 2] [3 0 0 2 2 1 1] [3 1 0 1 2 2 0] [3 1 0 1 2 0 2] [3 1 0 1 0 2 2] [3 1 0 2 1 2 0] [3 1 0 2 1 0 2] [3 1 0 2 2 1 0] [3 1 0 2 2 0 1] [3 1 0 2 0 2 1] [3 1 0 2 0 1 2] [3 1 0 0 2 2 1] [3 1 0 0 2 1 2] [3 1 0 0 1 2 2] [3 1 1 0 2 2 0] [3 1 1 0 2 0 2] [3 1 1 0 0 2 2] [3 1 1 2 0 2 0] [3 1 1 2 0 0 2] [3 1 1 2 2 0 0] [3 1 2 1 0 2 0] [3 1 2 1 0 0 2] [3 1 2 1 2 0 0] [3 1 2 0 1 2 0] [3 1 2 0 1 0 2] [3 1 2 0 2 1 0] [3 1 2 0 2 0 1] [3 1 2 0 0 2 1] [3 1 2 0 0 1 2] [3 1 2 2 0 1 0] [3 1 2 2 0 0 1] [3 1 2 2 1 0 0] [3 1 2 0 0 2 1] [3 1 2 0 0 1 2] [3 1 2 0 2 0 1] [3 1 2 0 2 1 0] [3 1 2 0 1 2 0] [3 1 2 0 1 0 2] [3 2 1 1 0 2 0] [3 2 1 1 0 0 2] [3 2 1 1 2 0 0] [3 2 1 0 1 2 0] [3 2 1 0 1 0 2] [3 2 1 0 2 1 0] [3 2 1 0 2 0 1] [3 2 1 0 0 2 1] [3 2 1 0 0 1 2] [3 2 1 2 0 1 0] [3 2 1 2 0 0 1] [3 2 1 2 1 0 0] [3 2 1 0 0 2 1] [3 2 1 0 0 1 2] [3 2 1 0 2 0 1] [3 2 1 0 2 1 0] [3 2 1 0 1 2 0] [3 2 1 0 1 0 2] [3 2 0 1 1 2 0] [3 2 0 1 1 0 2] [3 2 0 1 2 1 0] [3 2 0 1 2 0 1] [3 2 0 1 0 2 1] [3 2 0 1 0 1 2] [3 2 0 2 1 1 0] [3 2 0 2 1 0 1] [3 2 0 2 0 1 1] [3 2 0 0 1 2 1] [3 2 0 0 1 1 2] [3 2 0 0 2 1 1] [3 2 2 1 0 1 0] [3 2 2 1 0 0 1] [3 2 2 1 1 0 0] [3 2 2 0 1 1 0] [3 2 2 0 1 0 1] [3 2 2 0 0 1 1] [3 2 2 0 0 1 1] [3 2 2 0 1 0 1] [3 2 2 0 1 1 0] [3 2 0 1 0 2 1] [3 2 0 1 0 1 2] [3 2 0 1 2 0 1] [3 2 0 1 2 1 0] [3 2 0 1 1 2 0] [3 2 0 1 1 0 2] [3 2 0 0 1 2 1] [3 2 0 0 1 1 2] [3 2 0 0 2 1 1] [3 2 0 2 0 1 1] [3 2 0 2 1 0 1] [3 2 0 2 1 1 0]]

3 0 1 1 2 2 0
3 2 1 1 0 2 0

3 2 0 2 0 1 1
3 2 0 1 1 0 2

3 2 0 0 1 2 1
*/