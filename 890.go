package partice

import (
	"reflect"
)

func findAndReplacePattern(words []string, pattern string) []string {
	result:=make([]string,0)
	if len(words)==0||len(pattern)==0{
		return result
	}

	tp:=getPat(pattern)
	for _,v:=range words{
		if reflect.DeepEqual(tp,getPat(v)){
			result=append(result,v)
		}
	}

	return result
}

func getPat(p string) []int{
	m:=make(map[int32]int,0)
	res:=make([]int,len(p))
	for i,v:=range p{
		val,ok:=m[v]
		if ok{
			res[i]=val
		}else{
			res[i]=-1
		}
		m[v]=i
	}
	return res
}
