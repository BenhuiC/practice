package partice

import "testing"

func Test_kItemsWithMaximumSum(t *testing.T) {
	type args struct {
		numOnes    int
		numZeros   int
		numNegOnes int
		k          int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				numOnes:    3,
				numZeros:   2,
				numNegOnes: 0,
				k:          2,
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				numOnes:    3,
				numZeros:   2,
				numNegOnes: 0,
				k:          4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kItemsWithMaximumSum(tt.args.numOnes, tt.args.numZeros, tt.args.numNegOnes, tt.args.k); got != tt.want {
				t.Errorf("kItemsWithMaximumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
