package partice

import "testing"

func Test_reorderSpaces(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				text: "  this   is  a sentence ",
			},
			want: "this   is   a   sentence",
		},
		{
			name: "2",
			args: args{
				text: " practice   makes   perfect",
			},
			want: "practice   makes   perfect ",
		},
		{
			name: "3",
			args: args{
				text: "hello   world",
			},
			want: "hello   world",
		},
		{
			name: "4",
			args: args{
				text: "  walks  udp package   into  bar a",
			},
			want: "walks  udp  package  into  bar  a ",
		},
		{
			name: "5",
			args: args{
				text: "a",
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderSpaces(tt.args.text); got != tt.want {
				t.Errorf("reorderSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
