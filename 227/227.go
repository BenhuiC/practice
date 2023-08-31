package _27

import (
	"strconv"
)

func calculate(s string) int {
	if len(s)==0{
		return 0
	}
	rev:=parse(s)
	stack:=make([]int,0)
	for _,v:=range rev{
		switch tp:=v.(type) {
		case int:
			stack = append(stack, tp)
		case uint8:
			if len(stack)<2{
				break
			}
			v1:=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
			v2:=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
			var tmpCal int
			if tp=='-'{
				tmpCal=v2-v1
			}else if tp=='+'{
				tmpCal=v2+v1
			}else if tp=='*'{
				tmpCal=v2*v1
			}else{
				tmpCal=v2/v1
			}
			stack = append(stack, tmpCal)
		}
	}
	if len(stack)!=0{
		return stack[0]
	}
	return 0
}

func parse(s string) (da []interface{}){
	stk:=make([]byte,0)
	da=make([]interface{},0)
	for i:=0;i<len(s);{
		if s[i]>='0'&&s[i]<='9'{
			j:=i+1
			for ;j<len(s)&&s[j]>='0'&&s[j]<='9';j++{}
			v,_:=strconv.Atoi(s[i:j])
			da = append(da, v)
			i=j
			continue
		}
		if len(stk)==0{
			stk = append(stk, s[i])
			i++
			continue
		}
		if lv(stk[len(stk)-1])<lv(s[i]){
			stk = append(stk, s[i])
			i++
		}else{
			p:=stk[len(stk)-1]
			stk=stk[:len(stk)-1]
			da = append(da, p)
		}
	}
	for j:=len(stk)-1;j>=0;j--{
		da = append(da, stk[j])
	}
	return
}

// 符号等级
func lv(x byte) int{
	if x=='*'||x=='/'{
		return 2
	}else{
		return 1
	}
}