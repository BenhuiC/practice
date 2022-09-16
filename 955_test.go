package partice

import "testing"

func Test_minDeletionSize955(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				strs: []string{"ca", "bb", "ac"},
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				strs: []string{"xc", "yb", "za"},
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				strs: []string{"zyx", "wvu", "tsr"},
			},
			want: 3,
		},
		{
			name: "4",
			args: args{
				strs: []string{"aaa", "aaa", "aaa", "aaa"},
			},
			want: 0,
		},
		{
			name: "5",
			args: args{
				strs: []string{"abx", "agz", "bgc", "bfc"},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDeletionSize955(tt.args.strs); got != tt.want {
				t.Errorf("minDeletionSize955() = %v, want %v", got, tt.want)
			}
		})
	}
}
