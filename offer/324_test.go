package offer

import "testing"

func Test_wiggleSort(t *testing.T) {
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
				nums: []int{1, 3, 2, 1, 1, 1, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wiggleSort(tt.args.nums)
			t.Log(tt.args.nums)
		})
	}
}
