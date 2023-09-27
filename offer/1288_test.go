package offer

import "testing"

func Test_removeCoveredIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				intervals: [][]int{{1, 4}, {2, 8}, {3, 6}},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeCoveredIntervals(tt.args.intervals); got != tt.want {
				t.Errorf("removeCoveredIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
