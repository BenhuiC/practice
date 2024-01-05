package offer

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "1",
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'a',
			},
			want: 'c',
		},
		{
			name: "2",
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'c',
			},
			want: 'f',
		},
		{
			name: "3",
			args: args{
				letters: []byte{'x', 'x', 'y'},
				target:  'z',
			},
			want: 'x',
		},
		{
			name: "4",
			args: args{
				letters: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c', 'd', 'f', 'j'},
				target:  'b',
			},
			want: 'c',
		},
		{
			name: "5",
			args: args{
				letters: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c', 'd', 'f', 'j'},
				target:  'e',
			},
			want: 'f',
		},
		{
			name: "6",
			args: args{
				letters: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c', 'd', 'f', 'f', 'j'},
				target:  'e',
			},
			want: 'f',
		},
		{
			name: "7",
			args: args{
				letters: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c', 'd', 'f', 'j'},
				target:  'j',
			},
			want: 'a',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %c, want %c", got, tt.want)
			}
		})
	}
}
