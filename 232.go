package partice

type MyQueue struct {
	Value []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{Value: make([]int, 0)}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.Value = append(this.Value, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	var x int
	if len(this.Value) > 0 {
		x = this.Value[0]
		this.Value = this.Value[1:]
	}
	return x
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	var x int
	if len(this.Value) > 0 {
		x = this.Value[0]
	}
	return x
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.Value) <= 0
}
