package array

import "testing"

func Test_carPooling(t *testing.T) {
	type args struct {
		trips    [][]int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				trips:    [][]int{{2, 1, 1000}, {3, 3, 7}},
				capacity: 4,
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				trips:    [][]int{{2, 1, 5}, {3, 3, 7}},
				capacity: 5,
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				trips:    [][]int{{2, 1, 5}, {3, 5, 7}},
				capacity: 3,
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				trips:    [][]int{{9, 0, 1}, {3, 5, 7}},
				capacity: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := carPooling(tt.args.trips, tt.args.capacity); got != tt.want {
				t.Errorf("carPooling() = %v, want %v", got, tt.want)
			}
		})
	}
}
