package offer

import "testing"

func Test_minNumber(t *testing.T) {
	tests := []struct {
		args []int
	}{
		{
			args: []int{2, 3, 1, 0},
		},
		{
			args: []int{10, 11, 9},
		},
		{
			args: []int{128, 12},
		},
		{
			args: []int{3, 30, 34, 5, 9},
		},
		{
			args: []int{121, 12},
		},
		{
			args: []int{12, 1213},
		},
	}
	for _, tt := range tests {
		got := minNumber(tt.args)
		t.Log(got)
	}
}
