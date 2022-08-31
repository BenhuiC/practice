package partice

import (
	"testing"
)

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//{
		//	name: "1",
		//	args: args{
		//		num1: "2",
		//		num2: "3",
		//	},
		//	want: "6",
		//},
		//{
		//	name: "2",
		//	args: args{
		//		num1: "123",
		//		num2: "456",
		//	},
		//	want: "56088",
		//},
		//{
		//	name: "3",
		//	args: args{
		//		num1: "3414134",
		//		num2: "3141234",
		//	},
		//	want: "10724593801356",
		//},
		{
			name: "4",
			args: args{
				num1: "6",
				num2: "501",
			},
			want: "3006",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
