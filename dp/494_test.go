package dp

import "testing"

func Test_findTargetSumWays2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums:   []int{1, 1, 1, 1, 1},
				target: 3,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findTargetSumWays2() = %v, want %v", got, tt.want)
			}
		})
	}
}
