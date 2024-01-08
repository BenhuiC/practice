package offer

import "testing"

func Test_reachNumber(t *testing.T) {
	type args struct {
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				target: 999999,
			},
			want: 1414,
		},
		{
			name: "2",
			args: args{
				target: 2,
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				target: 3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reachNumber(tt.args.target); got != tt.want {
				t.Errorf("reachNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
