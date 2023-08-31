package _95

type MedianFinder struct {
	head *node
	len int
}

type node struct{
	val int
	next *node
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}


func (m *MedianFinder) AddNum(num int)  {
	if m.head==nil{
		m.head=&node{
			val:  num,
			next: nil,
		}
		m.len++
		return
	}
	if num<=m.head.val{
		tmp:=&node{
			val:  num,
			next: m.head,
		}
		m.head=tmp
		m.len++
		return
	}
	var p *node
	for p=m.head;p.next!=nil&&p.next.val<num;p=p.next{}
	p.next=&node{
		val:  num,
		next: p.next,
	}
	m.len++
}


func (m *MedianFinder) FindMedian() float64 {
	if m.len==1{
		return float64(m.head.val)
	}
	t:=m.len/2
	if m.len%2==0{
		t=t-1
	}
	var i int
	var p *node
	for p=m.head;i<t;{
		i++
		p=p.next
	}
	if m.len%2==1{
		return float64(p.val)
	}
	return float64(p.val+p.next.val)/2.0
}

