package partice

import "testing"

func Test_minimumRefill(t *testing.T) {
	type args struct {
		plants    []int
		capacityA int
		capacityB int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				plants:    []int{2, 2, 3, 3},
				capacityA: 5,
				capacityB: 5,
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				plants:    []int{2, 2, 3, 3},
				capacityA: 3,
				capacityB: 4,
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				plants:    []int{5},
				capacityA: 10,
				capacityB: 8,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumRefill(tt.args.plants, tt.args.capacityA, tt.args.capacityB); got != tt.want {
				t.Errorf("minimumRefill() = %v, want %v", got, tt.want)
			}
		})
	}
}
