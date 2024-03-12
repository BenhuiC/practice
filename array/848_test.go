package array

import "testing"

func Test_shiftingLetters(t *testing.T) {
	type args struct {
		s      string
		shifts []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s:      "abcjjjj",
				shifts: []int{3, 5, 9, 10, 12, 20, 100000},
			},
			want: "ljfdthn",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiftingLetters2(tt.args.s, tt.args.shifts); got != tt.want {
				t.Errorf("shiftingLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
