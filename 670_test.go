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
				num: 8736949,
			},
			want: 9736948,
		},
		{
			name: "2",
			args: args{
				num: 111,
			},
			want: 111,
		},
		{
			name: "3",
			args: args{
				num: 969,
			},
			want: 996,
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
