package partice

import "testing"

func Test_minCut(t *testing.T) {
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
			args: args{s: "eegiicgaeadbcfacfhifdbiehbgejcaeggcgbahfcajfhjjdgj"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCut(tt.args.s); got != tt.want {
				t.Errorf("minCut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isHui(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{s: "aa"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHui(tt.args.s); got != tt.want {
				t.Errorf("isHui() = %v, want %v", got, tt.want)
			}
		})
	}
}
