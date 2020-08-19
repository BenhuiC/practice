package partice

import (
	"reflect"
	"testing"
)

func Test_reverseList2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList2() = %v, want %v", got, tt.want)
			}
		})
	}
}
