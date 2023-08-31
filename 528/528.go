package _28

import "math/rand"

type Solution struct {
	data []int
	sv int
}


func Constructor(w []int) Solution {
	s:=0
	for i,v:=range w{
		s+=v
		w[i]=s
	}
	return Solution{
		data: w,
		sv:s,
	}
}


func (this *Solution) PickIndex() int {
	res:=0
	r:=rand.Intn(this.sv)
	for i:=len(this.data)-1;i>=0;i--{
		if this.data[i]>r{
			res=i
		}else{
			break
		}
	}
	return res
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
