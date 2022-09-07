package partice

import (
	"reflect"
	"testing"
)

func Test_prevPermOpt1(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				arr: []int{3, 2, 1},
			},
			want: []int{3, 1, 2},
		},
		{
			name: "2",
			args: args{
				arr: []int{1, 1, 5},
			},
			want: []int{1, 1, 5},
		},
		{
			name: "3",
			args: args{
				arr: []int{1, 9, 4, 6, 7},
			},
			want: []int{1, 7, 4, 6, 9},
		},
		{
			name: "4",
			args: args{
				arr: []int{3, 1, 1, 3},
			},
			want: []int{1, 3, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prevPermOpt1(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prevPermOpt1() = %v, want %v", got, tt.want)
			}
		})
	}
}
