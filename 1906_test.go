package partice

import (
	"reflect"
	"testing"
)

func Test_minDifference(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums:    []int{1, 3, 4, 8},
				queries: [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 3}},
			},
			want: []int{2, 1, 4, 1},
		},
		{
			name: "2",
			args: args{
				nums:    []int{4, 5, 2, 2, 7, 10},
				queries: [][]int{{2, 3}, {0, 2}, {0, 5}, {3, 5}},
			},
			want: []int{-1, 1, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDifference(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
2 2 4 5 7 10
2 3 0 1 4 5
2 3 -> -1
0 3 -> 1
0 1 -> 1
1 4 -> 2
4 5 -> 3
*/
