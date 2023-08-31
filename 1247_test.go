package partice

import "testing"

func Test_minimumSwap(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s1: "xx",
				s2: "yy",
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				s1: "xy",
				s2: "yx",
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				s1: "xx",
				s2: "yx",
			},
			want: -1,
		},
		{
			name: "4",
			args: args{
				s1: "xxyyxyxyxx",
				s2: "xyyxyxxxyx",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSwap(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("minimumSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
