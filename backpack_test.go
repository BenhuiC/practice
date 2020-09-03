package partice

import "testing"

func Test_backPack(t *testing.T) {
	type args struct {
		val []int
		c   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{val: []int{2, 3, 5, 7}, c: 11},
			want: 10,
		},
		{
			name: "2",
			args: args{val: []int{3, 4, 8, 5}, c: 10},
			want: 9,
		},
		{
			name: "3",
			args: args{val: []int{2, 3, 5, 7}, c: 12},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backPack2(tt.args.val, tt.args.c); got != tt.want {
				t.Errorf("backPack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backPack3(t *testing.T) {
	type args struct {
		val []int
		wei []int
		c   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				val: []int{1, 5, 2, 4},
				wei: []int{2, 3, 5, 7},
				c:   10,
			},
			want: 9,
		},
		{
			name: "2",
			args: args{
				val: []int{2, 5, 8},
				wei: []int{2, 3, 8},
				c:   10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backPack4(tt.args.val, tt.args.wei, tt.args.c); got != tt.want {
				t.Errorf("backPack3() = %v, want %v", got, tt.want)
			}
		})
	}
}
