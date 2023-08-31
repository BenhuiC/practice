package partice

import (
	"reflect"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				s: "1111",
			},
			want: []string{"1,1,1,1"},
		},
		{
			name: "2",
			args: args{
				s: "0000",
			},
			want: []string{"0,0,0,0"},
		},
		{
			name: "3",
			args: args{
				s: "11355",
			},
			want: []string{"0,0,0,0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddresses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
