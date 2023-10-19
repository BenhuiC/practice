package offer

type WordDictionary struct {
	root *treeNode
}

type treeNode struct {
	end   bool
	child map[byte]*treeNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &treeNode{child: make(map[byte]*treeNode)},
	}
}

func (this *WordDictionary) AddWord(word string) {
	r := this.root
	for i := range word {
		if r.child[word[i]] == nil {
			r.child[word[i]] = &treeNode{child: make(map[byte]*treeNode)}
		}
		r = r.child[word[i]]
	}
	r.end = true
}

func (this *WordDictionary) Search(word string) bool {
	var searchTree func(t *treeNode, w string) bool
	searchTree = func(t *treeNode, w string) bool {
		for i := range w {
			if w[i] >= 'a' && w[i] <= 'z' {
				t = t.child[w[i]]
				// 如果当前位置无对应字符，直接返回false
				if t == nil {
					return false
				}
				continue
			}
			// 处理 . 号
			for _, n := range t.child {
				// 跳过当前字符，直接判断子节点是否符合，有符合的就返回true
				if searchTree(n, w[i+1:]) {
					return true
				}
			}
			return false
		}
		return t.end == true
	}
	return searchTree(this.root, word)
}
