package partice

import (
	"testing"
)

func Test_constructArray(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				n: 3,
				k: 1,
			},
		},
		{
			name: "2",
			args: args{
				n: 3,
				k: 2,
			},
		},
		{
			name: "3",
			args: args{
				n: 100,
				k: 5,
			},
		},
		{
			name: "4",
			args: args{
				n: 5,
				k: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := constructArray(tt.args.n, tt.args.k)
			t.Log(got)
			m := map[int]struct{}{}
			for i := 1; i < len(got); i++ {
				m[Abs(got[i]-got[i-1])] = struct{}{}
			}
			if len(got) != tt.args.n || len(m) != tt.args.k {
				t.Errorf("constructArray() = %v,got len:%v, want %v", got, len(m), tt.args.k)
			}
		})
	}
}
