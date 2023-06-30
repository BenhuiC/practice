package partice

import (
	"reflect"
	"testing"
)

func Test_reconstructMatrix(t *testing.T) {
	type args struct {
		upper  int
		lower  int
		colsum []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				upper:  2,
				lower:  1,
				colsum: []int{1, 1, 1},
			},
			want: [][]int{{1, 1, 0}, {0, 0, 1}},
		},
		{
			name: "2",
			args: args{
				upper:  2,
				lower:  3,
				colsum: []int{2, 2, 1, 1},
			},
			want: [][]int{},
		},
		{
			name: "3",
			args: args{
				upper:  5,
				lower:  5,
				colsum: []int{2, 1, 2, 0, 1, 0, 1, 2, 0, 1},
			},
			want: [][]int{{1, 1, 1, 0, 1, 0, 0, 1, 0, 0}, {1, 0, 1, 0, 0, 0, 1, 1, 0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reconstructMatrix(tt.args.upper, tt.args.lower, tt.args.colsum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reconstructMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
