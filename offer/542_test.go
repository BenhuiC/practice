package offer

import (
	"reflect"
	"testing"
)

func Test_updateMatrix(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				mat: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}, {1, 0, 1}},
			},
			want: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}, {1, 0, 1}},
		},
		{
			name: "2",
			args: args{
				mat: [][]int{{0, 1, 0}, {0, 1, 0}, {0, 1, 0}, {0, 1, 0}, {0, 1, 0}},
			},
			want: [][]int{{0, 1, 0}, {0, 1, 0}, {0, 1, 0}, {0, 1, 0}, {0, 1, 0}},
		},
		{
			name: "3",
			args: args{
				mat: [][]int{{0, 1, 1, 0}},
			},
			want: [][]int{{0, 1, 1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateMatrix(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
