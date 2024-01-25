package offer

import "testing"

func Test_maxProfitAssignment(t *testing.T) {
	type args struct {
		difficulty []int
		profit     []int
		worker     []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				difficulty: []int{2, 4, 6, 8, 10, 10, 12, 13, 111, 222, 333},
				profit:     []int{10, 20, 30, 40, 50, 111, 234, 98, 9, 10, 12},
				worker:     []int{4, 5, 6, 70, 32, 21, 19, 11, 12, 14, 15, 5, 2, 123},
			},
			want: 2083,
		},
		{
			name: "2",
			args: args{
				difficulty: []int{2, 4, 6, 8, 10},
				profit:     []int{10, 20, 30, 40, 50},
				worker:     []int{4, 5, 6, 7},
			},
			want: 100,
		},
		{
			name: "3",
			args: args{
				difficulty: []int{2},
				profit:     []int{10},
				worker:     []int{1, 1},
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				difficulty: []int{2},
				profit:     []int{10},
				worker:     []int{2},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitAssignment(tt.args.difficulty, tt.args.profit, tt.args.worker); got != tt.want {
				t.Errorf("maxProfitAssignment() = %v, want %v", got, tt.want)
			}
		})
	}
}
