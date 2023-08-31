package partice

import "testing"

func Test_minOperations(t *testing.T) {
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
				nums: []int{1, 1, 1},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 5, 2, 4, 1},
			},
			want: 14,
		},
		{
			name: "3",
			args: args{
				nums: []int{8},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations1827(tt.args.nums); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
