package offer

import (
	"reflect"
	"testing"
)

func Test_findDiagonalOrder(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				mat: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			want: []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
		{
			name: "2",
			args: args{
				mat: [][]int{{1, 2, 3}, {4, 5, 6}},
			},
			want: []int{1, 2, 4, 5, 3, 6},
		},
		{
			name: "3",
			args: args{
				mat: [][]int{{1, 2, 3}},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "4",
			args: args{
				mat: [][]int{{1}, {2}, {3}},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "5",
			args: args{
				mat: [][]int{{1}},
			},
			want: []int{1},
		},
		{
			name: "6",
			args: args{
				mat: [][]int{{1, 2}, {3, 4}, {5, 6}},
			},
			want: []int{1, 2, 3, 5, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDiagonalOrder(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiagonalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
