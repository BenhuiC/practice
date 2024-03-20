package array

import "testing"

func Test_decodeAtIndex(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "a2345678999999999999999",
				k: 999,
			},
			want: "a",
		},
		{
			name: "2",
			args: args{
				s: "cpmxv8ewnfk3xxcilcmm68d2ygc88daomywc3imncfjgtwj8nrxjtwhiem5nzqnicxzo248g52y72v3yujqpvqcssrofd99lkovg",
				k: 480551547,
			},
			want: "x",
		},
		{
			name: "3",
			args: args{
				s: "leet2code3",
				k: 10,
			},
			want: "o",
		},
		{
			name: "4",
			args: args{
				s: "abc",
				k: 1,
			},
			want: "a",
		},
		{
			name: "5",
			args: args{
				s: "a23",
				k: 6,
			},
			want: "a",
		},
		{
			name: "6",
			args: args{
				s: "leet2code3",
				k: 8,
			},
			want: "t",
		},
		{
			name: "7",
			args: args{
				s: "h5xk8ar9s222v93y22w2",
				k: 311373,
			},
			want: "h",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeAtIndex(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("decodeAtIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
