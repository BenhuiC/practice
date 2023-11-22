package offer

import "testing"

func Test_totalHammingDistance(t *testing.T) {
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
				nums: []int{4, 14, 2},
			},
			want: 6,
		},
		{
			name: "2",
			args: args{
				nums: []int{4, 14, 2, 10},
			},
			want: 11,
		},
		{
			name: "3",
			args: args{
				nums: []int{4},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalHammingDistance(tt.args.nums); got != tt.want {
				t.Errorf("totalHammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
