package partice

import (
	"reflect"
	"testing"
)

func Test_prisonAfterNDays(t *testing.T) {
	type args struct {
		cells []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				cells: []int{0,1,0,1,1,0,0,1},
				n:     7,
			},
			want: []int{0,0,1,1,0,0,0,0},
		},
		{
			name: "2",
			args: args{
				cells: []int{1,0,0,1,0,0,1,0},
				n:     1000000000,
			},
			want: []int{0,0,1,1,1,1,1,0},
		},
		{
			name: "3",
			args: args{
				cells: []int{0,1,0,0,1,0,1,1},
				n:     1000,
			},
			want: []int{0,0,1,0,1,1,0,0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prisonAfterNDays(tt.args.cells, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prisonAfterNDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
