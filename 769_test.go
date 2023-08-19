package partice

import "testing"

func Test_maxChunksToSorted(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				arr: []int{4, 3, 2, 1, 0},
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				arr: []int{1, 0, 2, 3, 4},
			},
			want: 4,
		},
		{
			name: "3",
			args: args{
				arr: []int{1, 2, 0, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunksToSorted(tt.args.arr); got != tt.want {
				t.Errorf("maxChunksToSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
