package partice

import "testing"

func Test_totalFruit(t *testing.T) {
	type args struct {
		fruits []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				fruits: []int{1, 2, 1},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				fruits: []int{0, 1, 2, 2},
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				fruits: []int{1, 2, 3, 2, 2},
			},
			want: 4,
		},
		{
			name: "4",
			args: args{
				fruits: []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4},
			},
			want: 5,
		},
		{
			name: "5",
			args: args{
				fruits: []int{1, 0, 3, 4, 3},
			},
			want: 3,
		},
		{
			name: "6",
			args: args{
				fruits: []int{0, 1, 6, 6, 4, 4, 6},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalFruit(tt.args.fruits); got != tt.want {
				t.Errorf("totalFruit() = %v, want %v", got, tt.want)
			}
		})
	}
}
