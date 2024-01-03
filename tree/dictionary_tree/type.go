package dictionary_tree

type dictionaryTree struct {
	end      bool
	children map[byte]*dictionaryTree
}

func NewDictionaryTree() *dictionaryTree {
	return &dictionaryTree{
		children: make(map[byte]*dictionaryTree),
	}
}

func (d *dictionaryTree) Add(word string) {
	cur := d
	for i, w := range word {
		c, ok := cur.children[byte(w)]
		if !ok {
			c = &dictionaryTree{children: make(map[byte]*dictionaryTree)}
			cur.children[byte(w)] = c
		}
		if i == len(word)-1 {
			c.end = true
		}
		cur = c
	}
}
