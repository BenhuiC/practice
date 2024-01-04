package dp

import "testing"

func Test_deleteAndEarn(t *testing.T) {
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
				nums: []int{3, 4, 2},
			},
			want: 6,
		},
		{
			name: "2",
			args: args{
				nums: []int{2, 2, 3, 3, 3, 4},
			},
			want: 9,
		},
		{
			name: "3",
			args: args{
				nums: []int{1, 2, 2, 3, 3, 3, 4},
			},
			want: 10,
		},
		{
			name: "4",
			args: args{
				nums: []int{1, 2, 2, 3, 3, 3, 4, 64, 2, 3, 4, 5, 24, 3, 4, 33, 22, 25, 23, 41, 51, 423, 2, 2, 2, 34, 43, 4, 2, 56, 6, 432},
			},
			want: 1228,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteAndEarn(tt.args.nums); got != tt.want {
				t.Errorf("deleteAndEarn() = %v, want %v", got, tt.want)
			}
		})
	}
}
