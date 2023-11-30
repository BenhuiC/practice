package offer

import "testing"

func Test_checkSubarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				nums: []int{0},
				k:    1,
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				nums: []int{0, 1},
				k:    1,
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				nums: []int{0, 0, 1},
				k:    1,
			},
			want: true,
		},
		{
			name: "5",
			args: args{
				nums: []int{23, 2, 4, 6, 7, 2134, 22, 34, 5, 6, 7, 9, 0, 1, 909, 123, 68},
				k:    17,
			},
			want: true,
		},
		{
			name: "6",
			args: args{
				nums: []int{0, 1},
				k:    2,
			},
			want: false,
		},
		{
			name: "7",
			args: args{
				nums: []int{0, 1},
				k:    2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSubarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("checkSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
