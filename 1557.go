package partice


func findSmallestSetOfVertices(n int, edges [][]int) []int {
	result:=make([]int,0)
	from:=make(map[int]struct{})
	to:=make(map[int]struct{})

	for _,v:=range edges{
		from[v[0]]= struct{}{}
		to[v[1]]= struct{}{}
	}

	for k:=range from{
		if _,ok:=to[k];!ok{
			result = append(result, k)
		}
	}

	return result
}