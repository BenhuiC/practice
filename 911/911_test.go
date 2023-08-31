package _11

import "testing"

func TestTopVotedCandidate_Q(t *testing.T) {
	type fields struct {
		persons []int
		times   []int
	}
	type args struct {
		t int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "1",
			fields: fields{
				persons: []int{0, 0, 0, 0, 1},
				times:   []int{0, 6, 39, 52, 75},
			},
			args: args{
				t: 25,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(tt.fields.persons, tt.fields.times)
			if got := this.Q(tt.args.t); got != tt.want {
				t.Errorf("Q() = %v, want %v", got, tt.want)
			}
		})
	}
}
