package offer

import "testing"

func Test_cuttingRope(t *testing.T) {
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
			args: args{n: 10},
			want: 36,
		},
		{
			name: "2",
			args: args{n: 9},
			want: 27,
		},
		{
			name: "3",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "4",
			args: args{n: 15},
			want: 243,
		},
		{
			name: "5",
			args: args{n: 58},
			want: 1549681956,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cuttingRope(tt.args.n); got != tt.want {
				t.Errorf("cuttingRope() = %v, want %v", got, tt.want)
			}
		})
	}
}
