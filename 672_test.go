package partice

import "testing"

func Test_flipLights(t *testing.T) {
	type args struct {
		n       int
		presses int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n:       1,
				presses: 1,
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				n:       2,
				presses: 1,
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				n:       3,
				presses: 1,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipLights(tt.args.n, tt.args.presses); got != tt.want {
				t.Errorf("flipLights() = %v, want %v", got, tt.want)
			}
		})
	}
}
