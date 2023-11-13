package offer

import "testing"

func Test_compress(t *testing.T) {
	type args struct {
		chars []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				chars: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			},
			want: 6,
		},
		{
			name: "2",
			args: args{
				chars: []byte{'a'},
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				chars: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', '1', '1', 'A', 'A', '@', '@'},
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.args.chars); got != tt.want {
				t.Errorf("compress() = %v, want %v", got, tt.want)
			} else {
				t.Logf("got %v", string(tt.args.chars[:got]))
			}
		})
	}
}
