package dp

import "testing"

func Test_largestSumOfAverages(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1",
			args: args{
				nums: []int{9, 1, 2, 3, 9},
				k:    3,
			},
			want: 20.00000,
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    4,
			},
			want: 20.50000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSumOfAverages(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("largestSumOfAverages() = %v, want %v", got, tt.want)
			}
		})
	}
}
