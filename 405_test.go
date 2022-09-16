package partice

import "testing"

func Test_toHex(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				num: 160,
			},
			want: "a0",
		},
		{
			name: "-1",
			args: args{
				num: -1,
			},
			want: "ffffffff",
		},
		{
			name: "-1",
			args: args{
				num: -2,
			},
			want: "fffffffe",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toHex(tt.args.num); got != tt.want {
				t.Errorf("toHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
