package array

import "testing"

func Test_largestTimeFromDigits(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				arr: []int{1, 2, 3, 4},
			},
			want: "23:41",
		},
		{
			name: "2",
			args: args{
				arr: []int{5, 5, 5, 5},
			},
			want: "",
		},
		{
			name: "3",
			args: args{
				arr: []int{2, 4, 0, 0},
			},
			want: "20:40",
		},
		{
			name: "4",
			args: args{
				arr: []int{2, 5, 9, 9},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestTimeFromDigits(tt.args.arr); got != tt.want {
				t.Errorf("largestTimeFromDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
