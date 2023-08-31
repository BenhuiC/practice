package partice

import "testing"

func Test_removeDuplicates80(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 4},
			},
			want: 8,
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 4, 5, 6, 6, 7, 8, 8, 8},
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates80(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
