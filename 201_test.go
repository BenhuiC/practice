package partice

import "testing"

func Test_rangeBitwiseAnd(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				left:  5,
				right: 7,
			},
			want: 4,
		},
		{
			name: "2",
			args: args{
				left:  0,
				right: 0,
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				left:  1,
				right: 2147483647,
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				left:  0,
				right: 1,
			},
			want: 0,
		},
		{
			name: "5",
			args: args{
				left:  123,
				right: 123,
			},
			want: 123,
		},
		{
			name: "6",
			args: args{
				left:  0,
				right: 2,
			},
			want: 0,
		},
		{
			name: "7",
			args: args{
				left:  0,
				right: 5,
			},
			want: 0,
		},
		{
			name: "8",
			args: args{
				left:  2,
				right: 4,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeBitwiseAnd(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("rangeBitwiseAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
