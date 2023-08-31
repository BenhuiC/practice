package partice

import (
	"reflect"
	"testing"
)

func Test_peopleIndexes(t *testing.T) {
	type args struct {
		favoriteCompanies [][]string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				favoriteCompanies: [][]string{{"leetcode", "google", "facebook"}, {"google", "microsoft"}, {"google", "facebook"}, {"google"}, {"amazon"}},
			},
			want: []int{0, 1, 4},
		},
		{
			name: "2",
			args: args{
				favoriteCompanies: [][]string{{"nxaqhyoprhlhvhyojanr", "ovqdyfqmlpxapbjwtssm", "qmsbphxzmnvflrwyvxlc", "udfuxjdxkxwqnqvgjjsp", "yawoixzhsdkaaauramvg", "zycidpyopumzgdpamnty"}, {"nxaqhyoprhlhvhyojanr", "ovqdyfqmlpxapbjwtssm", "udfuxjdxkxwqnqvgjjsp", "yawoixzhsdkaaauramvg", "zycidpyopumzgdpamnty"}, {"ovqdyfqmlpxapbjwtssm", "qmsbphxzmnvflrwyvxlc", "udfuxjdxkxwqnqvgjjsp", "yawoixzhsdkaaauramvg", "zycidpyopumzgdpamnty"}},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peopleIndexes(tt.args.favoriteCompanies); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("peopleIndexes() = %v, want %v", got, tt.want)
			}
		})
	}
}
