package partice

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				digits: []int{},
			},
			want: []int{1},
		},
		{
			name: "2",
			args: args{
				digits: []int{0},
			},
			want: []int{1},
		},
		{
			name: "3",
			args: args{
				digits: []int{1},
			},
			want: []int{2},
		},
		{
			name: "4",
			args: args{
				digits: []int{1,2,3},
			},
			want: []int{1,2,4},
		},
		{
			name: "5",
			args: args{
				digits: []int{1,2,9},
			},
			want: []int{1,3,0},
		},
		{
			name: "6",
			args: args{
				digits: []int{1,9,9},
			},
			want: []int{2,0,0},
		},
		{
			name: "7",
			args: args{
				digits: []int{9,9,9},
			},
			want: []int{1,0,0,0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
