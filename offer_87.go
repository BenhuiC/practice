package partice

import (
	"strconv"
	"strings"
)

var res087 []string
var tmpStr []string

func restoreIpAddresses(s string) []string {
	res087=make([]string,0)
	tmpStr=make([]string,4)

	trySplit(s,0,0)
	return res087
}

func trySplit(s string,bg ,off int) {
	if off>=4||bg>=len(s){
		return
	}
	var ed int
	if s[bg]=='0'{
		ed=bg+1
	}else if s[bg]-'0'>2{
		ed=bg+2
	}else{
		ed=bg+3
	}
	for i:=bg+1;i<=ed&&i<=len(s);i++{
		tmp:=s[bg:i]
		v,_:=strconv.Atoi(tmp)
		if v>255{
			continue
		}
		tmpStr[off]=tmp
		if off==3&&i==len(s){
			res087=append(res087, strings.Join(tmpStr,"."))
			return
		}
		trySplit(s,i,off+1)
	}
}