package partice

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "3",
			args: args{str: "+134134"},
			want: 134134,
		},
		{
			name: "4",
			args: args{str: "-134134"},
			want: -134134,
		},
		{
			name: "5",
			args: args{str: "   134134werqj"},
			want: 134134,
		},
		{
			name: "6",
			args: args{str: "wewr134134"},
			want: 0,
		},
		{
			name: "7",
			args: args{str: "91283472332"},
			want: 2147483647,
		},
		{
			name: "8",
			args: args{str: "-91283472332"},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
