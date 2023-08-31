package partice

import "testing"

func Test_numberOfBoomerangs(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				points: [][]int{{1, 0}, {0, 0}, {2, 0}},
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				points: [][]int{{1, 1}, {0, 0}, {2, 2}},
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				points: [][]int{{1, 1}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfBoomerangs(tt.args.points); got != tt.want {
				t.Errorf("numberOfBoomerangs() = %v, want %v", got, tt.want)
			}
		})
	}
}
