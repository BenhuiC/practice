package partice

func prisonAfterNDays(cells []int, n int) []int {
	res:=make([]int,len(cells))
	if len(cells)<=2{
		return res
	}
	if cells[0]==1||cells[len(cells)-1]==1 {
		n = n - 1
		proc(cells)
	}
	n=n%(14)
	for i:=0;i<n;i++{
		proc(cells)
	}
	return cells
}

func proc(cells []int){
	tmp:=make([]int,len(cells))
	copy(tmp,cells)
	for i:=1;i<len(cells)-1;i++{
		if tmp[i-1]==tmp[i+1]{
			cells[i]=1
		}else{
			cells[i]=0
		}
	}
	cells[0]=0
	cells[len(cells)-1]=0

	return
}