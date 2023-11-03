package offer

import "testing"

func Test_validUtf8(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				data: []int{250, 145, 145, 145, 145},
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				data: []int{128},
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				data: []int{197, 130, 1, 197, 130, 1, 14, 197, 130, 1, 235, 140, 4},
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				data: []int{192, 128},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validUtf8(tt.args.data); got != tt.want {
				t.Errorf("validUtf8() = %v, want %v", got, tt.want)
			}
		})
	}
}
