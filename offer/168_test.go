package offer

import "testing"

func Test_convertToTitle(t *testing.T) {
	type args struct {
		columnNumber int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				columnNumber: 28,
			},
			want: "AB",
		},
		{
			name: "2",
			args: args{
				columnNumber: 2147483647,
			},
			want: "FXSHRXW",
		},
		{
			name: "3",
			args: args{
				columnNumber: 26,
			},
			want: "Z",
		},
		{
			name: "4",
			args: args{
				columnNumber: 676,
			},
			want: "YZ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.args.columnNumber); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
