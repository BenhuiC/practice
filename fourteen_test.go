package partice

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	var tests = []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{strs: []string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			name: "2",
			args: args{strs: []string{"c", "c"}},
			want: "c",
		},
		{
			name: "3",
			args: args{strs: []string{"cool", "coo", "coot"}},
			want: "coo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
