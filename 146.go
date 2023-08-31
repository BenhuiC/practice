package partice

type LRUCache struct {
	values map[int]*entry
	size   int
	cap    int
	header *entry
	tail   *entry
}

type entry struct {
	key   int
	value int
	next  *entry
	prev  *entry
}

func Constructor146(capacity int) LRUCache {
	h := &entry{}
	t := &entry{prev: h}
	h.next = t
	return LRUCache{
		values: map[int]*entry{},
		size:   0,
		cap:    capacity,
		header: h,
		tail:   t,
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.values[key]
	if !ok {
		return -1
	}

	this.moveToHead(v)

	return v.value
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.values[key]
	if ok {
		v.value = value
		this.moveToHead(v)
		return
	}
	this.size++
	n := &entry{value: value, key: key}
	this.values[key] = n
	// add to head
	n.next = this.header.next
	n.next.prev = n
	n.prev = this.header
	this.header.next = n

	// remove tail
	if this.size > this.cap {
		t := this.tail.prev
		t.prev.next = this.tail
		this.tail.prev = t.prev
		delete(this.values, t.key)
	}
}

func (this *LRUCache) moveToHead(v *entry) {
	v.next.prev = v.prev
	v.prev.next = v.next
	v.next = this.header.next
	v.next.prev = v
	this.header.next = v
	v.prev = this.header
}
