package offer

import "testing"

func Test_isIdealPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 0, 2, 3, 5, 4},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIdealPermutation(tt.args.nums); got != tt.want {
				t.Errorf("isIdealPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
