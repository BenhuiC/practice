package partice

type MinStack struct {
	min   int
	stack []int
}

/** initialize your data structure here. */
//func Constructor() MinStack {
//	return MinStack{stack: make([]int, 0), min: int(^uint(0) >> 1)}
//}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if x < this.min {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	p := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if p <= this.min {
		m := int(^uint(0) >> 1)
		for _, v := range this.stack {
			if v < m {
				m = v
			}
		}
		this.min = m
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
