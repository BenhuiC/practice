package partice

import (
	"sort"
)

// todo

type ipo struct {
	p int
	c int
}


func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	arys:=make([]*ipo,len(profits))
	for i:=0;i<len(profits);i++{
		arys[i]=&ipo{
			p: profits[i],
			c: capital[i],
		}
	}
	sort.Slice(arys, func(i, j int) bool {
		if arys[i].p==arys[j].p{
			return arys[i].c<arys[j].c
		}

		return arys[i].p<arys[j].p
	})

	isEnd:=false
	start:=0
	for j:=0;j<k;j++{
		// 从后往前
		if isEnd{
			for t:=start;t>=0;t--{
				if arys[t]!=nil&&arys[t].c<=w{
					w+=arys[t].p
					arys[t]=nil
					start=t-1
					break
				}
			}
		}else{
			// 从前往后
			maxt:=-1
			maxv:=0
			for t:=start;t<len(profits);t++{
				if arys[t]==nil{
					continue
				}
				if arys[t].c<=w{
					if arys[t].p>=maxv{
						maxt=t
						start=t
						maxv=arys[t].p
					}
					continue
				}
			}
			if maxt>=0{
				w+=maxv
				isEnd=maxt==len(profits)-1
				arys[maxt]=nil
			}
		}
	}
	return w
}
