package offer

import "testing"

func Test_findShortestSubArray(t *testing.T) {
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
				nums: []int{1, 2, 2, 4, 1},
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				nums: []int{2, 1, 1, 2, 1, 3, 3, 3, 1, 3, 1, 3, 2},
			},
			want: 7,
		},
		{
			name: "3",
			args: args{
				nums: []int{2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findShortestSubArray(tt.args.nums); got != tt.want {
				t.Errorf("findShortestSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
