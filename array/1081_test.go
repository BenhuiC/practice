package array

import "testing"

func Test_smallestSubsequence(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "bcabc",
			},
			want: "abc",
		},
		{
			name: "2",
			args: args{
				s: "cbacdcbc",
			},
			want: "acdb",
		},
		{
			name: "3",
			args: args{
				s: "cbaacabcaaccaacababa",
			},
			want: "abc",
		},
		{
			name: "4",
			args: args{
				s: "cbaacabcacb",
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestSubsequence(tt.args.s); got != tt.want {
				t.Errorf("smallestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
