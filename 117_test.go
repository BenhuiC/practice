package partice

import (
	"reflect"
	"testing"
)

func Test_connect(t *testing.T) {
	type args struct {
		root *Node117
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "1",
			args: args{
				root: &Node117{
					Val:  1,
					Left: &Node117{Val: -7},
					Right: &Node117{
						Val:   9,
						Left:  &Node117{Val: -1, Right: &Node117{Val: 8}},
						Right: &Node117{Val: -7, Left: &Node117{Val: -9}},
					},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connect(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
