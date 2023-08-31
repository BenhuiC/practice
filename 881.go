package partice

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	res881:=0
	for i,j:=0,len(people)-1;i<j;{
		if people[i]+people[j]>limit{
			j--
			continue
		}
		res881++
		people[i]=-1
		people[j]=-1
		i++
		j--
	}
	for _,v:=range people{
		if v!=-1{
			res881++
		}
	}
	return res881
}
