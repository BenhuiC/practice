package partice

import "testing"

func Test_minDeletionSize(t *testing.T) {
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
				strs: []string{"babca", "bbazb"},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				strs: []string{"edcba"},
			},
			want: 4,
		},
		{
			name: "3",
			args: args{
				strs: []string{"ghi", "def", "abc"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDeletionSize(tt.args.strs); got != tt.want {
				t.Errorf("minDeletionSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
