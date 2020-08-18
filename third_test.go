package partice

import (
	"fmt"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{s: "tmmzutx"},
			want: 5},
		{
			name: "2",
			args: args{s: "1112112"},
			want: 2},
		{
			name: "1",
			args: args{s: "bvbd"},
			want: 3},
		{
			name: "1",
			args: args{s: "pwwkew"},
			want: 3},
		{
			name: "3",
			args: args{s: "1"},
			want: 1},
		{
			name: "4",
			args: args{s: "abcdefghijklmnopqrstuvwxyza"},
			want: 26},
		{
			name: "5",
			args: args{s: ""},
			want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring3(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}

	fmt.Println(lengthOfLongestSubstring2("tmmzuxt"))
}
