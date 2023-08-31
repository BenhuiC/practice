package partice

import (
	"reflect"
	"testing"
)

func Test_advantageCount(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums1: []int{2, 0, 4, 1, 2},
				nums2: []int{1, 3, 0, 0, 2},
			},
			want: []int{2, 0, 1, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := advantageCount(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("advantageCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
