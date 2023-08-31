package partice

import (
	"reflect"
	"testing"
)

func Test_decrypt(t *testing.T) {
	type args struct {
		code []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				code: []int{2, 4, 9, 3},
				k:    -2,
			},
			want: []int{12, 5, 6, 13},
		},
		{
			name: "2",
			args: args{
				code: []int{10, 5, 7, 7, 3, 2, 10, 3, 6, 9, 1, 6},
				k:    -4,
			},
			want: []int{22, 26, 22, 28, 29, 22, 19, 22, 18, 21, 28, 19},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decrypt(tt.args.code, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
