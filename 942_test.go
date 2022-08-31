package partice

import (
	"reflect"
	"testing"
)

func Test_diStringMatch(t *testing.T) {
	tests := []struct {
		args string
		want []int
	}{
		{
			args: "III",
			want: []int{0, 1, 2, 3},
		},
		{
			args: "DDI",
			want: []int{3, 2, 0, 1},
		},
		{
			args: "IDID",
			want: []int{0, 4, 1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := diStringMatch(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
