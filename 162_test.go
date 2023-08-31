package partice

import "testing"

func Test_findPeakElement(t *testing.T) {
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
				nums: []int{3, 4, 3, 2, 1},
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 2, 1, 3, 5, 6, 4},
			},
			want: 5,
		},
		{
			name: "3",
			args: args{
				nums: []int{1},
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				nums: []int{1, 2},
			},
			want: 1,
		},
		{
			name: "5",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 2,
		},
		{
			name: "6",
			args: args{
				nums: []int{3, 2, 1},
			},
			want: 0,
		},
		{
			name: "7",
			args: args{
				nums: []int{3, 2, 1, 2, 3},
			},
			want: 0,
		},
		{
			name: "8",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6},
			},
			want: 5,
		},
		{
			name: "9",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakElement(tt.args.nums); got != tt.want {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
