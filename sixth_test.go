package partice

import (
	"fmt"
	"testing"
)

func Test_convert(t *testing.T) {

	s := "LEETCODESHIRING"
	fmt.Println(s[4:5])
	fmt.Println(s[7:8])
	fmt.Println(s[8:9])

	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s:       "LEETCODEISHIRING",
				numRows: 3,
			},
			want: "LCIRETOESIIGEDHN",
		},
		{
			name: "2",
			args: args{
				s:       "LEETCODEISHIRING",
				numRows: 4,
			},
			want: "LDREOEIIECIHNTSG",
		},
		{
			name: "3",
			args: args{
				s:       "LEETCODEISHIRING",
				numRows: 5,
			},
			want: "LIEESGEDHNTOIICR",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
