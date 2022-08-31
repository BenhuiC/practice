package partice

import (
	"reflect"
	"testing"
)

func Test_findClosestElements(t *testing.T) {
	type args struct {
		arr []int
		k   int
		x   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   3,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "2",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   5,
				x:   3,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "3",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   -1,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "4",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   -1,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "5",
			args: args{
				arr: []int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8},
				k:   3,
				x:   5,
			},
			want: []int{3, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findClosestElements(tt.args.arr, tt.args.k, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findClosestElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
