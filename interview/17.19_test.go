package interview

import (
	"reflect"
	"testing"
)

func Test_missingTwo(t *testing.T) {
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
				nums: []int{1},
			},
			want: []int{2, 3},
		},
		{
			name: "2",
			args: args{
				nums: []int{2, 3},
			},
			want: []int{4, 1},
		},
		{
			name: "3",
			args: args{
				nums: []int{2, 3, 4, 5},
			},
			want: []int{6, 1},
		},
		{
			name: "4",
			args: args{
				nums: []int{2},
			},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingTwo(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("missingTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
