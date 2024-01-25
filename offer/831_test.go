package offer

import "testing"

func Test_maskPII(t *testing.T) {
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
				s: "LeetCode@LeetCode.com",
			},
			want: "l*****e@leetcode.com",
		},
		{
			name: "2",
			args: args{
				s: "AB@qq.com",
			},
			want: "a*****b@qq.com",
		},
		{
			name: "3",
			args: args{
				s: "1(234)567-890",
			},
			want: "***-***-7890",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maskPII(tt.args.s); got != tt.want {
				t.Errorf("maskPII() = %v, want %v", got, tt.want)
			}
		})
	}
}
