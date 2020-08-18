package partice

import (
	"fmt"
	"testing"
)

func Test_generateTree(t *testing.T) {
	type args struct {
		arg []string
		n   int
	}
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "1",
			args: []string{"-1", "0", "3", "-2", "4", "null", "null", "8"},
		},
	}
	for _, tt := range tests {
		fmt.Println("----------------")
		pre(generateTree(tt.args, 0))
		fmt.Println("\n----------------")
	}
}
