package partice

import "testing"

func Test_maximumSum(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				arr: []int{1, -2, 0, 3},
			},
			want: 4,
		},
		{
			name: "2",
			args: args{
				arr: []int{1, -2, -2, 3},
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				arr: []int{-1, -1, -1, -1},
			},
			want: -1,
		},
		{
			name: "4",
			args: args{
				arr: []int{2, 1, -2, -5, -2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSum2(tt.args.arr); got != tt.want {
				t.Errorf("maximumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
