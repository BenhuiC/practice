package partice

import "testing"

func Test_numTimesAllBlue(t *testing.T) {
	type args struct {
		flips []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				flips: []int{3, 2, 4, 1, 5},
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				flips: []int{4, 1, 2, 3},
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				flips: []int{1, 2, 3},
			},
			want: 3,
		},
		{
			name: "4",
			args: args{
				flips: []int{1, 3, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTimesAllBlue(tt.args.flips); got != tt.want {
				t.Errorf("numTimesAllBlue() = %v, want %v", got, tt.want)
			}
		})
	}
}
