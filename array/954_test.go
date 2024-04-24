package array

import "testing"

func Test_canReorderDoubled(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				arr: []int{0, 0, 0, 1, 1, 1},
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				arr: []int{4, -2, 2, -4},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canReorderDoubled(tt.args.arr); got != tt.want {
				t.Errorf("canReorderDoubled() = %v, want %v", got, tt.want)
			}
		})
	}
}
