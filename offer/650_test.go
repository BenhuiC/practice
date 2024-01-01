package offer

import "testing"

func Test_minSteps(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				n: 1,
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				n: 5,
			},
			want: 5,
		},
		{
			name: "4",
			args: args{
				n: 999,
			},
			want: 46,
		},
		{
			name: "5",
			args: args{
				n: 27,
			},
			want: 9,
		},
		{
			name: "6",
			args: args{
				n: 1000,
			},
			want: 21,
		},
		{
			name: "7",
			args: args{
				n: 189,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSteps(tt.args.n); got != tt.want {
				t.Errorf("minSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
