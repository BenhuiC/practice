package partice

import "testing"

func Test_mincostTickets(t *testing.T) {
	type args struct {
		days  []int
		costs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				days:  []int{1, 4, 6, 7, 8, 20},
				costs: []int{2, 7, 15},
			},
			want: 11,
		},
		{
			name: "2",
			args: args{
				days:  []int{1, 4, 6, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22, 23, 27, 28},
				costs: []int{3, 13, 45},
			},
			want: 44,
			//
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mincostTickets2(tt.args.days, tt.args.costs); got != tt.want {
				t.Errorf("mincostTickets() = %v, want %v", got, tt.want)
			}
		})
	}
}
