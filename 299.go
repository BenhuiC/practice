package partice

import "fmt"

func getHint(secret string, guess string) string {
	m:=make(map[int32]int,0)
	for _,v:=range secret{
		m[v-'0']++
	}
	var a,b int
	for i,v:=range guess{
		if secret[i]==guess[i]{
			a++
		}
		if m[v-'0']>0{
			b++
			m[v-'0']--
		}
	}
	b=b-a
	return fmt.Sprintf("%dA%dB",a,b)
}
