package array

import "testing"

func Test_subarraysDivByK(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{-1, 2, 9},
				k:    2,
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				nums: []int{2, -2, 2, -4},
				k:    6,
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				nums: []int{5, 5, 0, -2, -3, 1},
				k:    5,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraysDivByK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraysDivByK() = %v, want %v", got, tt.want)
			}
		})
	}
}
