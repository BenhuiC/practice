package partice

func canCompleteCircuit(gas []int, cost []int) int {
	res134:=-1
	for start:=0;start<len(gas);{
		if cost[start]<=gas[start]{
			last:=0
			i:=start
			for ;i<start+len(gas);i++{
				tIndex:=i%len(gas)
				last+=gas[tIndex]
				last-=cost[tIndex]
				if last>=0{
					continue
				}
				start=i+1
				break
			}
			if last>=0&&i==start+len(gas){
				res134=start
				break
			}
		}else{
			start++
		}
	}

	return res134
}