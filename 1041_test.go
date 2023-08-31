package partice

import "testing"

func Test_isRobotBounded(t *testing.T) {
	type args struct {
		instructions string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				instructions: "GGLLGG",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				instructions: "GG",
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				instructions: "GL",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRobotBounded(tt.args.instructions); got != tt.want {
				t.Errorf("isRobotBounded() = %v, want %v", got, tt.want)
			}
		})
	}
}
