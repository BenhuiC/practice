package partice

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				path: "/home////",
			},
			want: "/home",
		},
		{
			name: "2",
			args: args{
				path: "/",
			},
			want: "/",
		},
		{
			name: "3",
			args: args{
				path: "/../",
			},
			want: "/",
		},
		{
			name: "4",
			args: args{
				path: "/home////../.././c",
			},
			want: "/c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
