package partice

import "testing"

func Test_maxResult(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1,-1,-2,4,-7,3},
				k:    2,
			},
			want: 7,
		},
		{
			name: "2",
			args: args{
				nums: []int{10,-5,-2,4,0,3},
				k:    2,
			},
			want: 15,
		},
		{
			name: "3",
			args: args{
				nums: []int{10,-5,-2,4,0,3},
				k:    3,
			},
			want: 17,
		},
		{
			name: "4",
			args: args{
				nums: []int{1,-5,-20,4,-1,3,-6,-3},
				k:    2,
			},
			want: 0,
		},
		{
			name: "5",
			args: args{
				nums: []int{1,2},
				k:    2,
			},
			want: 3,
		},
		{
			name: "6",
			args: args{
				nums: []int{1},
				k:    2,
			},
			want: 1,
		},
		{
			name: "7",
			args: args{
				nums: []int{100,-100,-300,-300,-300,-100,100},
				k:    4,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxResult(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
