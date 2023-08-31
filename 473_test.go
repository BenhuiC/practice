package partice

import "testing"

func Test_makesquare(t *testing.T) {
	type args struct {
		matchsticks []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				matchsticks: []int{1, 1, 2, 2, 2},
			},
			want: true,
		},
		{
			name: "1",
			args: args{
				matchsticks: []int{3, 3, 3, 3, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makesquare(tt.args.matchsticks); got != tt.want {
				t.Errorf("makesquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
