package dp

import "testing"

func Test_canPartition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 5, 5, 11},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 2, 3, 5},
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				nums: []int{2, 2, 3, 5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition2(tt.args.nums); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
