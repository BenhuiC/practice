package partice

import "testing"

func Test_maxScoreSightseeingPair(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				values: []int{8, 1, 5, 2, 6},
			},
			want: 11,
		},
		{
			name: "2",
			args: args{
				values: []int{1, 2, 2, 3, 3, 3, 4, 5, 5, 1, 4, 10, 34, 123, 34, 21, 3, 1, 3, 4, 52, 5, 5, 6, 6, 2},
			},
			want: 168,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScoreSightseeingPair(tt.args.values); got != tt.want {
				t.Errorf("maxScoreSightseeingPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
