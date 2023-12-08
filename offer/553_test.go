package offer

import (
	"testing"
)

func Test_optimalDivision(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				nums: []int{1000, 100, 10, 2},
			},
			want: "1000/(100/10/2)",
		},
		{
			name: "2",
			args: args{
				nums: []int{2, 3, 4},
			},
			want: "2/(3/4)",
		},
		{
			name: "3",
			args: args{
				nums: []int{1000, 100, 10, 2, 123, 1000, 999, 11, 23, 39},
			},
			want: "1000/(100/10/2/123/1000/999/11/23/39)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := optimalDivision(tt.args.nums); got != tt.want {
				t.Errorf("optimalDivision() = %v, want %v", got, tt.want)
			}
		})
	}
}
