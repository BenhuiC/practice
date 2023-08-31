package _11

type WordDictionary struct {
	words map[uint8]*WordDictionary
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		words: map[uint8]*WordDictionary{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	if word == "" {
		this.words[0] = &WordDictionary{}
		return
	}
	if _, ok := this.words[word[0]]; !ok {
		this.words[word[0]] = &WordDictionary{words: map[uint8]*WordDictionary{}}
	}
	this.words[word[0]].AddWord(word[1:])
}

func (this *WordDictionary) Search(word string) bool {
	if word == "" {
		_, ok := this.words[0]
		return ok
	}
	if _, ok := this.words[word[0]]; !ok {
		return false
	}

}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
