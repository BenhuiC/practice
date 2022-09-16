package partice

import "testing"

func Test_reorderedPowerOf2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				n: 123,
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				n: 1,
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				n: 2,
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				n: 134217728,
			},
			want: true,
		},
		{
			name: "5",
			args: args{
				n: 134271728,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderedPowerOf2(tt.args.n); got != tt.want {
				t.Errorf("reorderedPowerOf2() = %v, want %v", got, tt.want)
			}
		})
	}
}
