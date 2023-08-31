package partice

import "testing"

func Test_evaluate(t *testing.T) {
	type args struct {
		s         string
		knowledge [][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s:         "(name)is(age)yearsold",
				knowledge: [][]string{{"name","bob"},{"age","10"}},
			},
			want: "bobis10yearsold",
		},
		{
			name: "2",
			args: args{
				s:         "(name)is(age)",
				knowledge: [][]string{{"name","bob"}},
			},
			want: "bobis?",
		},
		{
			name: "3",
			args: args{
				s:         "(name)is(name)yearsold",
				knowledge: [][]string{{"name","bob"},{"age","10"}},
			},
			want: "bobisbobyearsold",
		},
		{
			name: "4",
			args: args{
				s:         "hi(name)",
				knowledge: [][]string{{"age","10"}},
			},
			want: "hi?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluate(tt.args.s, tt.args.knowledge); got != tt.want {
				t.Errorf("evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
