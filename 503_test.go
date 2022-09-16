package partice

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElements(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 1},
			},
			want: []int{2, -1, 2},
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 2, 3, 4, 3},
			},
			want: []int{2, 3, 4, -1, 4},
		},
		{
			name: "3",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: []int{2, 3, 4, 5, -1},
		},
		{
			name: "4",
			args: args{
				nums: []int{-3, -2, -1, 0, 1},
			},
			want: []int{-2, -1, 0, 1, -1},
		},
		{
			name: "5",
			args: args{
				nums: []int{0, -2, -3},
			},
			want: []int{-1, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElements(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
