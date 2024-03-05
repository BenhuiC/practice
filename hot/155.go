package hot

type MinStack struct {
	min   int
	stack []int
}

func Constructor155() MinStack {
	return MinStack{stack: make([]int, 0), min: int(^uint(0) >> 1)}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, 0)
		this.min = x
		return
	}
	diff := x - this.min
	this.stack = append(this.stack, diff)
	if diff < 0 {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	diff := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	if diff < 0 {
		this.min = this.min - diff
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	diff := this.stack[len(this.stack)-1]
	if diff < 0 {
		return this.min
	} else {
		return this.min + diff
	}
}

func (this *MinStack) GetMin() int {
	return this.min
}
