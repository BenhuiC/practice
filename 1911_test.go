package partice

import "testing"

func Test_maxAlternatingSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "1",
			args: args{
				nums: []int{4, 2, 5, 3},
			},
			want: 7,
		},
		{
			name: "2",
			args: args{
				nums: []int{5, 6, 7, 8},
			},
			want: 8,
		},
		{
			name: "3",
			args: args{
				nums: []int{6, 2, 1, 2, 4, 5},
			},
			want: 10,
		},
		{
			name: "4",
			args: args{
				nums: []int{5, 4, 3},
			},
			want: 5,
		},
		{
			name: "5",
			args: args{
				nums: []int{10},
			},
			want: 10,
		},
		{
			name: "6",
			args: args{
				nums: []int{2, 1, 5, 4, 4},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAlternatingSum(tt.args.nums); got != tt.want {
				t.Errorf("maxAlternatingSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
