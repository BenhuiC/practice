package offer

import "testing"

func Test_pushDominoes(t *testing.T) {
	type args struct {
		dominoes string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				dominoes: ".L.R...LR..L..",
			},
			want: "LL.RR.LLRRLL..",
		},
		{
			name: "2",
			args: args{
				dominoes: "RLRLR...L......R................R..............L.....RR....RR",
			},
			want: "RLRLRR.LL......RRRRRRRRRRRRRRRRRRRRRRRRRLLLLLLLL.....RRRRRRRR",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pushDominoes(tt.args.dominoes); got != tt.want {
				t.Errorf("pushDominoes() = %v, want %v", got, tt.want)
			}
		})
	}
}
