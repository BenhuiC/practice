package partice

import "testing"

func Test_removeKdigits(t *testing.T) {
	type args struct {
		num string
		k   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				num: "1432219",
				k:   1,
			},
			want: "132219",
		},
		{
			name: "2",
			args: args{
				num: "52660469",
				k:   2,
			},
			want: "260469",
		},
		{
			name: "3",
			args: args{
				num: "112",
				k:   1,
			},
			want: "11",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeKdigits(tt.args.num, tt.args.k); got != tt.want {
				t.Errorf("removeKdigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
