package partice

import "testing"

func Test_monotoneIncreasingDigits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n: 1234321,
			},
			want: 1233999,
		},
		{
			name: "2",
			args: args{
				n: 1234,
			},
			want: 1234,
		},
		{
			name: "3",
			args: args{
				n: 20,
			},
			want: 19,
		},
		{
			name: "4",
			args: args{
				n: 10,
			},
			want: 9,
		},
		{
			name: "5",
			args: args{
				n: 668841,
			},
			want: 667999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := monotoneIncreasingDigits(tt.args.n); got != tt.want {
				t.Errorf("monotoneIncreasingDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
