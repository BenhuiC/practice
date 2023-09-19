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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permuteUnique(tt.args.nums)
			t.Log(got)
		})
	}
}
