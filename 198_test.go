package partice

import "testing"

func Test_rob(t *testing.T) {
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
				nums: []int{1,2,3,1},
			},
			want: 4,
		},
		{
			name: "2",
			args: args{
				nums: []int{2,7,9,3,1},
			},
			want: 12,
		},
		{
			name: "3",
			args: args{
				nums: []int{3,2,19,20},
			},
			want: 23,
		},
		{
			name: "4",
			args: args{
				nums: []int{6,7,8,2,7,9,3,1},
			},
			want: 24,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
