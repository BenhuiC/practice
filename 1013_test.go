package partice

import "testing"

func Test_canThreePartsEqualSum(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				arr: []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				arr: []int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1},
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				arr: []int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canThreePartsEqualSum(tt.args.arr); got != tt.want {
				t.Errorf("canThreePartsEqualSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
