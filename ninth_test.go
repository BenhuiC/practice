package partice

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{x: 1},
			want: true,
		},
		{
			name: "2",
			args: args{x: 123},
			want: false,
		},
		{
			name: "3",
			args: args{x: 121},
			want: true,
		},
		{
			name: "4",
			args: args{x: 12321},
			want: true,
		},
		{
			name: "5",
			args: args{x: 111111},
			want: true,
		},
		{
			name: "6",
			args: args{x: 12222},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
