package partice

import (
	"reflect"
	"testing"
)

func Test_allPathsSourceTarget(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				graph: [][]int{{1, 2}, {3}, {3}, {}},
			},
			want: [][]int{{0,1,3},{0,2,3}},
		},
		{
			name: "2",
			args: args{
				graph: [][]int{{4,3,1}, {3,2,4}, {3}, {4},{}},
			},
			want: [][]int{{0,4},{0,3,4},{0,1,4},{0,1,3,4},{0,1,2,3,4}},
		},
		{
			name: "3",
			args: args{
				graph: [][]int{{1}, {}},
			},
			want: [][]int{{0,1}},
		},
		{
			name: "4",
			args: args{
				graph: [][]int{{3,1}, {4,5,7,2,5},{4,6,3},{6,4},{7,6,5},{6},{7},{}},
			},
			want: [][]int{{0,1}},
		},
		{
			name: "5",
			args: args{
				graph: [][]int{{2}, {},{1}},
			},
			want: [][]int{{0,2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPathsSourceTarget(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPathsSourceTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
