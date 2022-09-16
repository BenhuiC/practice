package partice

import "testing"

func Test_numDecodings(t *testing.T) {
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
				s: "12",
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				s: "226",
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				s: "0",
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				s: "06",
			},
			want: 0,
		},
		{
			name: "5",
			args: args{
				s: "12345",
			},
			want: 3,
		},
		{
			name: "6",
			args: args{
				s: "10",
			},
			want: 1,
		},
		{
			name: "7",
			args: args{
				s: "207",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
