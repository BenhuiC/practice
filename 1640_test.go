package partice

import "testing"

func Test_canFormArray(t *testing.T) {
	type args struct {
		arr    []int
		pieces [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				arr:    []int{15, 88},
				pieces: [][]int{{88}, {15}},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				arr:    []int{49, 18, 16},
				pieces: [][]int{{16, 18, 49}},
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				arr:    []int{91, 4, 64, 78},
				pieces: [][]int{{78}, {4, 64}, {91}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFormArray(tt.args.arr, tt.args.pieces); got != tt.want {
				t.Errorf("canFormArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
