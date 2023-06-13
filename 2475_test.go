package partice

import "testing"

func Test_unequalTriplets(t *testing.T) {
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
				nums: []int{2, 3, 4, 4, 4, 5},
			},
			want: 10,
		},

		{
			name: "2",
			args: args{
				nums: []int{1, 1, 1, 1},
			},
			want: 0,
		},

		{
			name: "3",
			args: args{
				nums: []int{4, 4, 2, 4, 3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unequalTriplets(tt.args.nums); got != tt.want {
				t.Errorf("unequalTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
