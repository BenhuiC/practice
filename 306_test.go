package partice

import "testing"

func Test_isAdditiveNumber(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				num: "1235813",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				num: "199100199",
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				num: "000",
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				num: "1023",
			},
			want: false,
		},
		{
			name: "5",
			args: args{
				num: "199111992",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAdditiveNumber(tt.args.num); got != tt.want {
				t.Errorf("isAdditiveNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
