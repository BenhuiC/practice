package partice

func allPathsSourceTarget(graph [][]int) [][]int {
	e:=len(graph)
	res797:=make([][]int,0)
	tp:=make([][]int,0)
	tp = append(tp, []int{0})
	for {
		if len(tp)==0{
			break
		}
		for i:=0;i<len(tp);i++{
			pp:=tp[0]
			tp=tp[1:]
			tgt :=graph[pp[len(pp)-1]]
			if pp[len(pp)-1]==e-1{
				res797 = append(res797, pp)
				continue
			}
			for _,v:=range tgt{
				p:=make([]int,len(pp)+1,len(pp)+1)
				copy(p,pp)
				p[len(p)-1]=v
				tp = append(tp, p)
			}
		}

	}
	return res797
}
