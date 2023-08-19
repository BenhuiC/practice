package partice

import "testing"

func Test_numSteps(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s: "1101",
			},
			want: 6,
		},
		{
			name: "2",
			args: args{
				s: "10",
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				s: "1",
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				s: "1101011111101101010101010011",
			},
			want: 39,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSteps(tt.args.s); got != tt.want {
				t.Errorf("numSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
