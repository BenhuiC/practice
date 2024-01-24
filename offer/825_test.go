package offer

import "testing"

func Test_numFriendRequests(t *testing.T) {
	type args struct {
		ages []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				ages: []int{20, 30, 100, 110, 120},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				ages: []int{16, 16},
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				ages: []int{16, 17, 18},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numFriendRequests(tt.args.ages); got != tt.want {
				t.Errorf("numFriendRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}
