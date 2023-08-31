package partice

func predictPartyVictory(senate string) string {
	rs:=make([]int,0)
	ds:=make([]int,0)
	for i,v:=range senate{
		if v=='R'{
			rs=append(rs,i)
		}else{
			ds=append(ds,i)
		}
	}

	n:=len(senate)
	for {
		if len(rs)==0{
			return "Dire"
		}
		if len(ds)==0{
			return "Radiant"
		}

		if rs[0]<ds[0]{
			rs=append(rs,rs[0]+n)
		}else{
			ds=append(ds,ds[0]+n)
		}

		rs=rs[1:]
		ds=ds[1:]
	}
}
