package partice

import "testing"

func Test_maximumSwap(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				num: 9909,
			},
			want: 9990,
		},
		{
			name: "2",
			args: args{
				num: 1234,
			},
			want: 4231,
		},
		{
			name: "3",
			args: args{
				num: 1,
			},
			want: 1,
		},
		{
			name: "4",
			args: args{
				num: 124443,
			},
			want: 424413,
		},
		{
			name: "5",
			args: args{
				num: 4321,
			},
			want: 4321,
		},
		{
			name: "6",
			args: args{
				num: 959,
			},
			want: 995,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSwap(tt.args.num); got != tt.want {
				t.Errorf("maximumSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
