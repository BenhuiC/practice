package offer

import "testing"

func Test_canTransform(t *testing.T) {
	type args struct {
		start string
		end   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				start: "RXXLRXRXL",
				end:   "XRLXXRRLX",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canTransform(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("canTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}
