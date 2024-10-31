package array

import "testing"

func Test_fractionAddition(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{"-1/2+1/2"},
			want: "0/1",
		},
		{
			name: "2",
			args: args{"-1/2+1/2+1/3"},
			want: "1/3",
		},
		{
			name: "3",
			args: args{"7/2+2/3-3/4"},
			want: "41/12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fractionAddition(tt.args.expression); got != tt.want {
				t.Errorf("fractionAddition() = %v, want %v", got, tt.want)
			}
		})
	}
}
