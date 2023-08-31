package partice

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	if version1==version2{
		return 0
	}
	res165:=0
	ary1:=strings.Split(version1,".")
	ary2:=strings.Split(version2,".")
	var i int
	for i=0;i<len(ary1)&&i<len(ary2);i++{
		v1,_:=strconv.ParseInt(ary1[i],10,64)
		v2,_:=strconv.ParseInt(ary2[i],10,64)
		if v1>v2{
			return 1
		}else if v1<v2{
			return -1
		}
	}
	if len(ary1)==len(ary2){
		return 0
	}

	if i==len(ary1){
		for j:=i;j<len(ary2);j++{
			v,_:=strconv.ParseInt(ary2[j],10,64)
			if v>0{
				return -1
			}
		}
	}else{
		for j:=i;j<len(ary1);j++{
			v,_:=strconv.ParseInt(ary1[j],10,64)
			if v>0{
				return 1
			}
		}
	}


	return res165
}
