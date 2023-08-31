package partice

import "fmt"

func evaluate(s string, knowledge [][]string) string {
	if s==""{
		return ""
	}
	var res string
	m:=make(map[string]string,0)
	for _,v:=range knowledge{
		m[v[0]]=v[1]
	}

	for i:=0;i<len(s);i++{
		if s[i]=='('{
			var j int
			for j=i;j<len(s)&&s[j]!=')';j++{}
			k,ok:=m[s[i+1:j]]
			if ok{
				res+=k
			}else{
				res+="?"
			}
			i=j
		}else{
			res+=fmt.Sprintf("%c",s[i])
		}
	}
	return res
}
