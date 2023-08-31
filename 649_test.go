package partice

import "testing"

func Test_predictPartyVictory(t *testing.T) {
	type args struct {
		senate string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				senate: "RRDDD",
			},
			want: "Radiant",
		},
		{
			name: "2",
			args: args{
				senate: "RD",
			},
			want: "Radiant",
		},
		{
			name: "3",
			args: args{
				senate: "RRDD",
			},
			want: "Radiant",
		},
		{
			name: "4",
			args: args{
				senate: "RRRDDRD",
			},
			want: "Radiant",
		},
		{
			name: "5",
			args: args{
				senate: "RDDDRRD",
			},
			want: "Dire",
		},
		{
			name: "6",
			args: args{
				senate: "RDDR",
			},
			want: "Radiant",
		},
		{
			name: "7",
			args: args{
				senate: "RRDRDDRDRRDDDDDRDRDR",
			},
			want: "Radiant",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := predictPartyVictory(tt.args.senate); got != tt.want {
				t.Errorf("predictPartyVictory() = %v, want %v", got, tt.want)
			}
		})
	}
}
