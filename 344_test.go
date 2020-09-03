package partice

import (
	"fmt"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{s: []byte{'h', 'e', 'l'}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			fmt.Println(tt.args.s)
		})
	}
}
