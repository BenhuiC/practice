package offer

import "testing"

func Test_soupServings(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1",
			args: args{
				n: 50,
			},
			want: 0.625,
		},
		{
			name: "2",
			args: args{
				n: 100,
			},
			want: 0.71875,
		},
		{
			name: "3",
			args: args{
				n: 999,
			},
			want: 0.97657,
		},
		{
			name: "4",
			args: args{
				n: 800,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := soupServings(tt.args.n); got != tt.want {
				t.Errorf("soupServings() = %v, want %v", got, tt.want)
			}
		})
	}
}
