package offer

import "testing"

func Test_expressiveWords(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s:     "heeellooo",
				words: []string{"hello", "hi", "helo"},
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				s:     "helllllooo",
				words: []string{"hello", "hel", "helo"},
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				s:     "dddiiiinnssssssoooo",
				words: []string{"dinnssoo", "ddinso", "ddiinnso", "ddiinnssoo", "ddiinso", "dinsoo", "ddiinsso", "dinssoo", "dinso"},
			},
			want: 3,
		},
		{
			name: "4",
			args: args{
				s:     "abbb",
				words: []string{"abb"},
			},
			want: 1,
		},
		{
			name: "5",
			args: args{
				s:     "aaa",
				words: []string{"aaaa"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expressiveWords(tt.args.s, tt.args.words); got != tt.want {
				t.Errorf("expressiveWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
