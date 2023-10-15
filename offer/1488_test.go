package offer

import (
	"reflect"
	"testing"
)

func Test_avoidFlood(t *testing.T) {
	type args struct {
		rains []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				rains: []int{1, 2, 3, 4},
			},
			want: []int{-1, -1, -1, -1},
		},
		{
			name: "2",
			args: args{
				rains: []int{1, 2, 0, 0, 2, 1},
			},
			want: []int{-1, -1, 2, 1, -1, -1},
		},
		{
			name: "3",
			args: args{
				rains: []int{1, 2, 0, 1, 2},
			},
			want: nil,
		},
		{
			name: "4",
			args: args{
				rains: []int{1, 2, 0, 1, 0, 1, 2},
			},
			want: nil,
		},
		{
			name: "5",
			args: args{
				rains: []int{0, 0, 3, 3},
			},
			want: nil,
		},
		{
			name: "6",
			args: args{
				rains: []int{1, 0, 2, 0, 3, 0, 2, 0, 0, 0, 1, 2, 3},
			},
			want: []int{-1, 1, -1, 2, -1, 3, -1, 2, 1, 1, -1, -1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := avoidFlood(tt.args.rains); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("avoidFlood() = %v, want %v", got, tt.want)
			}
		})
	}
}
