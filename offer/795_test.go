package offer

import "testing"

func Test_numSubarrayBoundedMax(t *testing.T) {
	type args struct {
		nums  []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums:  []int{2, 1, 4, 3},
				left:  2,
				right: 3,
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				nums:  []int{16, 69, 88, 85, 79, 87, 37, 33, 39, 34},
				left:  55,
				right: 57,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarrayBoundedMax2(tt.args.nums, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("numSubarrayBoundedMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
