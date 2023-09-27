package offer

import (
	"reflect"
	"testing"
)

func Test_diffWaysToCompute(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				expression: "2-1-1",
			},
			want: []int{0, 2},
		},
		{
			name: "2",
			args: args{
				expression: "2*3-4*5",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diffWaysToCompute(tt.args.expression); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diffWaysToCompute() = %v, want %v", got, tt.want)
			}
		})
	}
}
