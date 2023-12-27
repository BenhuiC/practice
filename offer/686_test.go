package offer

import "testing"

func Test_repeatedStringMatch(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				a: "abcd",
				b: "cdabcdab",
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				a: "abc",
				b: "xyzd",
			},
			want: -1,
		},
		{
			name: "3",
			args: args{
				a: "aaac",
				b: "aac",
			},
			want: 1,
		},
		{
			name: "4",
			args: args{
				a: "abccb",
				b: "cbabccb",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedStringMatch(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("repeatedStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
