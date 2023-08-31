package partice

import (
	"reflect"
	"testing"
)

func Test_numsSameConsecDiff(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				n: 3,
				k: 7,
			},
			want: []int{181, 292, 707, 818, 929},
		},
		{
			name: "2",
			args: args{
				n: 2,
				k: 0,
			},
			want: []int{11, 22, 33, 44, 55, 66, 77, 88, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numsSameConsecDiff(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numsSameConsecDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
