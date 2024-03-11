package array

import "testing"

func Test_longestMountain(t *testing.T) {
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
				arr: []int{2, 1, 4, 7, 3, 2, 5},
			},
			want: 5,
		},
		{
			name: "2",
			args: args{
				arr: []int{2, 2, 2},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestMountain(tt.args.arr); got != tt.want {
				t.Errorf("longestMountain() = %v, want %v", got, tt.want)
			}
		})
	}
}
