package offer_59

type MaxQueue struct {
	max []int
	queue []int
}


func Constructor() MaxQueue {
	return MaxQueue{
		max:   []int{},
		queue: []int{},
	}
}


func (this *MaxQueue) Max_value() int {
	if len(this.max)==0{
		return -1
	}
	return this.max[0]
}


func (this *MaxQueue) Push_back(value int)  {
	this.queue = append(this.queue, value)
	if len(this.max)==0||value<this.max[len(this.max)-1]{
		this.max = append(this.max, value)
		return
	}
	if value>this.max[0]{
		this.max=[]int{value}
		return
	}
	for j:=len(this.max)-1;j>=0&&this.max[j]<value;j--{
		this.max=this.max[:len(this.max)-1]
	}
	this.max = append(this.max, value)
}


func (this *MaxQueue) Pop_front() int {
	if len(this.queue)==0{
		return -1
	}
	r:=this.queue[0]
	this.queue=this.queue[1:]
	if this.max[0]==r{
		this.max=this.max[1:]
	}
	return r
}


/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
