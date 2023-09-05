package types

type node struct {
	key   int
	value int
	pre   *node
	next  *node
}

type LRUCache struct {
	head *node
	tail *node
	m    map[int]*node
	cap  int
}

func Constructor(capacity int) LRUCache {
	h := &node{}
	return LRUCache{
		head: h,
		tail: h,
		m:    make(map[int]*node),
		cap:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.m[key]
	if !ok {
		return -1
	}
	if n != this.head.next {
		n.pre.next = n.next
		if n.next != nil {
			n.next.pre = n.pre
		} else {
			this.tail = n.pre
		}
		n.pre = this.head
		n.next = this.head.next
		this.head.next.pre = n
		this.head.next = n
	}
	return n.value
}

func (this *LRUCache) Put(key int, value int) {
	val := this.Get(key)
	if val == -1 {
		// 写入头部
		n := &node{key: key, value: value}
		n.pre = this.head
		n.next = this.head.next
		this.head.next = n
		if n.next != nil {
			n.next.pre = n
		} else {
			this.tail = n
		}
		this.m[key] = n
		// 删除尾部
		if len(this.m) > this.cap {
			t := this.tail
			this.tail = t.pre
			this.tail.next = nil
			t.pre = nil
			delete(this.m, t.key)
		}
	} else {
		// 如果key存在，则对应node已移动到头部，直接修改头部的value
		this.head.next.value = value
	}
}
