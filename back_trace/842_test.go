package back_trace

import (
	"reflect"
	"testing"
)

func Test_splitIntoFibonacci(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				num: "12358132134",
			},
			want: []int{1, 2, 3, 5, 8, 13, 21, 34},
		},
		{
			name: "2",
			args: args{
				num: "1101111",
			},
			want: []int{11, 0, 11, 11},
		},
		{
			name: "3",
			args: args{
				num: "0123",
			},
			want: nil,
		},
		{
			name: "4",
			args: args{
				num: "539834657215398346785398346991079669377161950407626991734534318677529701785098211336528511",
			},
			want: nil,
		},
		{
			name: "5",
			args: args{
				num: "0000",
			},
			want: []int{0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitIntoFibonacci(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitIntoFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
539834657,21,539834678,539834699,1079669377,1619504076,2699173453,4318677529,7017850982,11336528511
*/
