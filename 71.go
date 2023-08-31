package partice

import "strings"

func simplifyPath(path string) string {
	if len(path)<=1{
		return path
	}
	res71:=""
	ary:=strings.Split(path,"/")
	if len(ary)==0{
		return path
	}
	tmp:=make([]string,0,len(ary))
	for _,v:=range ary{
		if v=="."||v==""{
			continue
		}
		if v==".."{
			if len(tmp)!=0 {
				tmp = tmp[:len(tmp)-1]
			}
			continue
		}
		tmp=append(tmp,v)
	}
	res71="/"+strings.Join(tmp,"/")

	return res71
}