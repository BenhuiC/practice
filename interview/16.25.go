package interview

type LRUCache struct {
	head *dulNode
	tail *dulNode
	len  int
	cap  int
}

type dulNode struct {
	key  int
	val  int
	next *dulNode
	prev *dulNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{}
	l.cap = capacity
	return l
}

func (this *LRUCache) Get(key int) int {
	if this.len == 0 {
		return -1
	}
	res := -1
	var tgt *dulNode
	for tgt = this.head; tgt != nil; tgt = tgt.next {
		if tgt.key == key {
			res = tgt.val
			break
		}
	}
	if tgt == nil || tgt == this.head {
		return res
	}
	// 前节点
	tgt.prev.next = tgt.next
	// 后节点
	if tgt != this.tail {
		tgt.next.prev = tgt.prev
	} else {
		this.tail = this.tail.prev
	}
	tgt.prev = nil
	tgt.next = nil

	// 放至头部
	tgt.next = this.head
	this.head.prev = tgt
	this.head = tgt

	return res
}

func (this *LRUCache) Put(key int, value int) {
	if this.Get(key) != -1 {
		this.head.val = value
		return
	}
	t := &dulNode{
		key: key,
		val: value,
	}
	if this.head == nil {
		this.head = t
		this.tail = this.head
		this.len++
		return
	}
	this.head.prev = t
	t.next = this.head
	this.head = t
	this.len++
	if this.len > this.cap {
		this.tail = this.tail.prev
		this.tail.next.prev = nil
		this.tail.next = nil
		this.len = this.cap
	}
}
