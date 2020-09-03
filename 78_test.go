package partice

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{nums: []int{1, 2, 3}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				fmt.Println(result)
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
