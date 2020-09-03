package partice

import "testing"

func Test_wordBreak2(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{s: "leetcode", wordDict: []string{"leet", "code"}},
			want: true,
		},
		{
			name: "2",
			args: args{s: "aaaaaaa", wordDict: []string{"aaa", "aaaa"}},
			want: true,
		},
		{
			name: "3",
			args: args{s: "codedoge", wordDict: []string{"code", "codedog", "doge"}},
			want: true,
		},
		{
			name: "4",
			args: args{s: "catsandog", wordDict: []string{"cats", "cat", "sand", "and"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak2(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreak2() = %v, want %v", got, tt.want)
			}
		})
	}
}
