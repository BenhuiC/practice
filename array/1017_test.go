package array

import "testing"

func Test_baseNeg2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				n: 199,
			},
			want: "111011011",
		},
		{
			name: "2",
			args: args{
				n: 2,
			},
			want: "110",
		},
		{
			name: "3",
			args: args{
				n: 6,
			},
			want: "11010",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := baseNeg2(tt.args.n); got != tt.want {
				t.Errorf("baseNeg2() = %v, want %v", got, tt.want)
			}
		})
	}
}
