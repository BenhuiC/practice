package array

import (
	"reflect"
	"testing"
)

func Test_powerfulIntegers(t *testing.T) {
	type args struct {
		x     int
		y     int
		bound int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				x:     2,
				y:     3,
				bound: 10,
			},
			want: []int{2, 4, 10, 3, 5, 7, 9},
		},
		{
			name: "2",
			args: args{
				x:     1,
				y:     2,
				bound: 10,
			},
			want: []int{2, 3, 5, 9},
		},
		{
			name: "3",
			args: args{
				x:     2,
				y:     1,
				bound: 10,
			},
			want: []int{2, 3, 5, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := powerfulIntegers(tt.args.x, tt.args.y, tt.args.bound); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("powerfulIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
