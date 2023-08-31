package _27

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s: "1+2+3",
			},
			want: 6,
		},
		{
			name: "2",
			args: args{
				s: "1*2+3",
			},
			want: 5,
		},
		{
			name: "3",
			args: args{
				s: "1*2+3*4",
			},
			want: 14,
		},
		{
			name: "4",
			args: args{
				s: "1/2+3*4",
			},
			want: 12,
		},
		{
			name: "5",
			args: args{
				s: "1*4",
			},
			want: 4,
		},
		{
			name: "6",
			args: args{
				s: "3/2",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
