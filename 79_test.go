package partice

import "testing"

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				board: [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}},
				word:  "ABCCED",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				board: [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}},
				word:  "SEE",
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				board: [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}},
				word:  "ABCB",
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				board: [][]byte{{'A','A','A','A'},{'A','A','A','A'},{'A','A','A','A'}},
				word:  "AAAAAAAAAAAAA",
			},
			want: false,
		},
		{
			name: "5",
			args: args{
				board: [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}},
				word:  "SES",
			},
			want: false,
		},
		{
			name: "6",
			args: args{
				board: [][]byte{{'C','A','A'},{'A','A','A'},{'B','C','D'}},
				word:  "AAB",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}