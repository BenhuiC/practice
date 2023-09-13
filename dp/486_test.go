package dp

import "testing"

func Test_predictTheWinner(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 5, 2},
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 5, 233, 7},
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				nums: []int{1, 2, 233, 7, 23},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := predictTheWinner(tt.args.nums); got != tt.want {
				t.Errorf("predictTheWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
