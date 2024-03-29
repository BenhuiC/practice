package partice

import "testing"

func Test_maxSubarraySumCircular(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, -2, 3, -2},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				nums: []int{5,-3,5},
			},
			want: 10,
		},
		{
			name: "3",
			args: args{
				nums: []int{3,-1,2,-1},
			},
			want: 4,
		},
		{
			name: "4",
			args: args{
				nums: []int{3},
			},
			want: 3,
		},
		{
			name: "5",
			args: args{
				nums: []int{-2,-3,-1},
			},
			want: -1,
		},
		{
			name: "6",
			args: args{
				nums: []int{3,-1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarraySumCircular(tt.args.nums); got != tt.want {
				t.Errorf("maxSubarraySumCircular() = %v, want %v", got, tt.want)
			}
		})
	}
}
