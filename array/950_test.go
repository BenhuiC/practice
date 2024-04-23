package array

import (
	"reflect"
	"testing"
)

func Test_deckRevealedIncreasing(t *testing.T) {
	type args struct {
		deck []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				deck: []int{17, 13, 11, 2, 3, 5, 7},
			},
			want: []int{2, 13, 3, 11, 5, 17, 7},
		},
		{
			name: "2",
			args: args{
				deck: []int{17, 13, 11, 2, 3, 5, 7, 19},
			},
			want: []int{2, 11, 3, 17, 5, 13, 7, 19},
		},
		{
			name: "3",
			args: args{
				deck: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: []int{1, 9, 2, 6, 3, 8, 4, 7, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deckRevealedIncreasing(tt.args.deck); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deckRevealedIncreasing() = %v, want %v", got, tt.want)
			}
		})
	}
}
