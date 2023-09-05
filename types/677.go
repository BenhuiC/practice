package types

type MapSum struct {
	data *entity
}

type entity struct {
	d    byte
	n    int
	sum  int
	next [26]*entity
}

func ConstructorMapSum() MapSum {
	return MapSum{data: &entity{next: [26]*entity{}}}
}

func (this *MapSum) Insert(key string, val int) {
	e := this.data
	for _, v := range key {
		if e.next[v-'a'] == nil {
			e.next[v-'a'] = &entity{next: [26]*entity{}, d: byte(v)}
		}
		e = e.next[v-'a']
	}
	diff := val - e.n
	e.n = val

	e = this.data
	for _, v := range key {
		e.next[v-'a'].sum += diff
		e = e.next[v-'a']
	}
}

func (this *MapSum) Sum(prefix string) int {
	e := this.data
	for _, v := range prefix {
		if e.next[v-'a'] == nil {
			return 0
		}
		e = e.next[v-'a']
	}
	return e.sum
}
