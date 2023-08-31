package partice

import (
	"reflect"
	"testing"
)

func Test_partitionLabels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				s: "ababcbacadefegdehijhklij",
			},
			want: []int{9, 7, 8},
		},
		{
			name: "2",
			args: args{
				s: "aa",
			},
			want: []int{2},
		},
		{
			name: "3",
			args: args{
				s: "ab",
			},
			want: []int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionLabels(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
