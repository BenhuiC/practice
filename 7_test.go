package partice

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				x: 123,
			},
			want: 321,
		},
		{
			name: "2",
			args: args{
				x: 0,
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				x: -123,
			},
			want: -321,
		},
		{
			name: "4",
			args: args{
				x: 1234567890,
			},
			want: 987654321,
		},
		{
			name: "5",
			args: args{
				x: 214748361543,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse2(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
