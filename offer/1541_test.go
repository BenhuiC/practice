package offer

import "testing"

func Test_minInsertions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s: "(()))",
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				s: "())",
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				s: "))())(",
			},
			want: 3,
		},
		{
			name: "4",
			args: args{
				s: "(()))(()))()())))",
			},
			want: 4,
		},
		{
			name: "5",
			args: args{
				s: ")))))))",
			},
			want: 5,
		},
		{
			name: "6",
			args: args{
				s: "(((()(()((())))(((()())))()())))(((()(()()((()()))",
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minInsertions(tt.args.s); got != tt.want {
				t.Errorf("minInsertions() = %v, want %v", got, tt.want)
			}
		})
	}
}
