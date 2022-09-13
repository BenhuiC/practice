package partice

import "testing"

func Test_minSteps(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s: "bab",
				t: "aba",
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				s: "leetcode",
				t: "practice",
			},
			/*
				l:1
				e:2
				t:0
				c:-1
				o:1
				d:1
				p:-1
				r:-1
				a:-1
				i:-1
			*/
			want: 5,
		},
		{
			name: "3",
			args: args{
				s: "anagram",
				t: "mangaar",
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				s: "xxyyzz",
				t: "xxyyzz",
			},
			want: 0,
		},
		{
			name: "5",
			args: args{
				s: "friend",
				t: "family",
			},
			/*
				f:0
				r:1
				i:0
				e:1
				n:1
				d:1
				a:-1
				m:-1
				l:-1
				y:-1
			*/
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSteps(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
