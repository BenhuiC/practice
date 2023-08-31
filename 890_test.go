package partice

import (
	"reflect"
	"testing"
)

func Test_findAndReplacePattern(t *testing.T) {
	type args struct {
		words   []string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				words:   []string{"abc","deq","mee","aqq","dkd","ccc"},
				pattern: "aaa",
			},
			want: []string{"ccc"},
		},
		{
			name: "2",
			args: args{
				words:   []string{"abc","deq","mee","aqq","dkd","ccc"},
				pattern: "acd",
			},
			want: []string{"abc","deq"},
		},
		{
			name: "3",
			args: args{
				words:   []string{"abc","deq","mee","aqq","dkd","ccc"},
				pattern: "nvn",
			},
			want: []string{"dkd"},
		},
		{
			name: "4",
			args: args{
				words:   []string{"abc","deq","mee","aqq","dkd","ccc"},
				pattern: "aee",
			},
			want: []string{"mee","aqq"},
		},
		{
			name: "5",
			args: args{
				words:   []string{"abc","deq","mee","aqq","dkd","ccc"},
				pattern: "qqb",
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAndReplacePattern(tt.args.words, tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAndReplacePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
