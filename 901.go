package partice

type StockSpanner struct {
	stack []int
	m     map[int]item
	idx   int
}

type item struct {
	idx, nu int
}

func Constructor901() StockSpanner {
	return StockSpanner{
		stack: make([]int, 0),
		m:     make(map[int]item),
	}
}

func (this *StockSpanner) Next(price int) int {
	this.idx++
	if len(this.stack) == 0 {
		this.stack = append(this.stack, price)
		this.m[price] = item{idx: this.idx, nu: 1}
		return 1
	}
	for len(this.stack) > 1 && this.stack[len(this.stack)-2] <= price {
		delete(this.m, this.stack[len(this.stack)-1])
		this.stack = this.stack[:len(this.stack)-1]
	}

	if this.stack[len(this.stack)-1] <= price {
		t := this.m[this.stack[len(this.stack)-1]]
		res := t.nu + this.idx - t.idx

		delete(this.m, this.stack[len(this.stack)-1])
		this.stack[len(this.stack)-1] = price
		this.m[price] = item{idx: this.idx, nu: res}
		return res
	}
	this.stack = append(this.stack, price)
	this.m[price] = item{idx: this.idx, nu: 1}

	return 1
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
