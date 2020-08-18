package partice

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "yes",
			args: args{s: "aaaaa"},
			want: "aaaaa",
		},
		{
			name: "yes2",
			args: args{s: "aabaa"},
			want: "aabaa",
		},
		{
			name: "yes3",
			args: args{s: "aaa"},
			want: "aaa",
		},
		{
			name: "yes4",
			args: args{s: "aaaba"},
			want: "aaa",
		},
		{
			name: "yes5",
			args: args{s: "abaabad"},
			want: "abaaba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
