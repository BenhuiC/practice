package partice

import "strings"

func reverseWords(s string) string {
	ary:=strings.Split(strings.Trim(s," ")," ")
	res151:=""
	tmp:=make([]string,0)
	for i:=len(ary)-1;i>=0;i--{
		if ary[i]==""{
			continue
		}
		tmp=append(tmp,ary[i])
	}
	res151=strings.Join(tmp," ")
	return res151
}
