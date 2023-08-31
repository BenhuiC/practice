package partice

import (
	"fmt"
	"strconv"
	"strings"
)

func isAdditiveNumber(num string) bool {
	nlen:=len(num)
	if nlen<=2{
		return false
	}
	for i:=1;i<=nlen/2;i++{
		if i-0!=1&&strings.HasPrefix(num[:i],"0"){
			continue
		}
		for j:=i+1;j<nlen;j++{
			if j-i!=1&&strings.HasPrefix(num[i:j],"0"){
				continue
			}
			for k:=j+1;k<=nlen;k++{
				if k-j!=1&&strings.HasPrefix(num[j:k],"0"){
					continue
				}
				v1,_:=strconv.Atoi(num[:i])
				v2,_:=strconv.Atoi(num[i:j])
				v3,_:=strconv.Atoi(num[j:k])
				if v1+v2==v3&&findv(num,nlen,v2,v3,k){
					return true
				}
			}
		}
	}

	return false
}

func findv(num string,nl,v2,v3,st int) bool{
	for {
		if st==nl{
			return true
		}
		vtp:=v2+v3
		vtpSt:=fmt.Sprint(vtp)
		if st+len(vtpSt)>nl{
			return false
		}
		if vtpSt!=num[st:st+len(vtpSt)]{
			return false
		}
		v2,v3,st=v3,vtp,st+len(vtpSt)
	}
}