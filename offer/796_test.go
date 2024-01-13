package offer

import "testing"

func Test_rotateString(t *testing.T) {
	type args struct {
		s    string
		goal string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s:    "gcmbf",
				goal: "fgcmb",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateString(tt.args.s, tt.args.goal); got != tt.want {
				t.Errorf("rotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
