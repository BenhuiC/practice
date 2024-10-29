package offer

import (
	"reflect"
	"testing"
)

func Test_validStrings(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{3},
			want: []string{"010", "011", "101", "110", "111"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validStrings(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
