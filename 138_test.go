package partice

import (
	"reflect"
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	ns := []*Node{{Val: 7}, {Val: 13}, {Val: 11}, {Val: 10}, {Val: 1}}
	ns[0].Next = ns[1]
	ns[1].Next = ns[2]
	ns[2].Next = ns[3]
	ns[3].Next = ns[4]
	ns[0].Random = nil
	ns[1].Random = ns[0]
	ns[2].Random = ns[4]
	ns[3].Random = ns[2]
	ns[4].Random = ns[0]

	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "1",
			args: args{head: ns[0]},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyRandomList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyRandomList() = %v, want %v", got, tt.want)
			}
		})
	}
}
