package hot

type Trie struct {
	ent *entry
}

type entry struct {
	val  byte
	end  bool
	next [26]*entry
}

func Constructor() Trie {
	return Trie{
		ent: &entry{next: [26]*entry{}},
	}
}

func (this *Trie) Insert(word string) {
	e := this.ent
	for _, v := range word {
		if e.next[v-'a'] == nil {
			e.next[v-'a'] = &entry{
				val:  byte(v),
				next: [26]*entry{},
			}
		}
		e = e.next[v-'a']
	}
	e.end = true
}

func (this *Trie) Search(word string) bool {
	e := this.ent
	for _, v := range word {
		if e.next[v-'a'] == nil {
			return false
		}
		e = e.next[v-'a']
	}

	return e.end
}

func (this *Trie) StartsWith(prefix string) bool {
	e := this.ent
	for _, v := range prefix {
		if e.next[v-'a'] == nil {
			return false
		}
		e = e.next[v-'a']
	}
	return true
}
