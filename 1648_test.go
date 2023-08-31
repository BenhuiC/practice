package partice

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		inventory []int
		orders    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				inventory: []int{2,5},
				orders:    4,
			},
			want: 14,
		},
		{
			name: "2",
			args: args{
				inventory: []int{3,5},
				orders:    6,
			},
			want: 19,
		},
		{
			name: "3",
			args: args{
				inventory: []int{2,4,6,8,10},
				orders:    20,
			},
			want: 110,
		},
		{
			name: "4",
			args: args{
				inventory: []int{1000000000},
				orders:    1000000000,
			},
			want: 500000000500000000,
		},
		{
			name: "5",
			args: args{
				inventory: []int{10,10,10,10,10},
				orders:    10,
			},
			want: 95,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.inventory, tt.args.orders); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
