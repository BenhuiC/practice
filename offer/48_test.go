package offer

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "1",
			args: "pwwkew",
			want: 3,
		},
		{
			name: "2",
			args: "bbbbb",
			want: 1,
		},
		{
			name: "3",
			args: "abcabcbb",
			want: 3,
		},
		{
			name: "4",
			args: "a",
			want: 1,
		},
		{
			name: "5",
			args: "aa",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
