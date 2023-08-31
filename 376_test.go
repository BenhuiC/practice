package partice

import "testing"

func Test_wiggleMaxLength(t *testing.T) {
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
				nums: []int{1, 7, 4, 9, 2, 5},
			},
			want: 6,
		},
		{
			name: "2",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				nums: []int{1, 1},
			},
			want: 1,
		},
		{
			name: "4",
			args: args{
				nums: []int{1, 7},
			},
			want: 2,
		},
		{
			name: "5",
			args: args{
				nums: []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8},
			},
			want: 7,
		},
		{
			name: "6",
			args: args{
				nums: []int{3, 3, 3, 2, 5},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wiggleMaxLength(tt.args.nums); got != tt.want {
				t.Errorf("wiggleMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
