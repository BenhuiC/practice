package partice

import "testing"

func Test_maxTurbulenceSize(t *testing.T) {
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
				arr: []int{9, 4, 2, 10, 7, 8, 8, 1, 9},
			},
			want: 5,
		},
		{
			name: "2",
			args: args{
				arr: []int{4, 8, 12, 16},
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				arr: []int{100},
			},
			want: 1,
		},
		{
			name: "4",
			args: args{
				arr: []int{1, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTurbulenceSize(tt.args.arr); got != tt.want {
				t.Errorf("maxTurbulenceSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
