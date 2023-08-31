package partice

type Trie struct {
	V []*t
}

type t struct{
	fg bool
	next []*t
}


/** Initialize your data structure here. */
//func Constructor() Trie {
//	return Trie{V: make([]*t,26)}
//}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	if len(word)==0{
		return
	}
	ary:=this.V
	for i:=0;i<len(word);i++{
		n:=word[i]-'a'
		if ary[n]==nil{
			ary[n]=&t{next: make([]*t,26)}
		}
		if i==len(word)-1{
			ary[n].fg=true
		}
		ary=ary[n].next
	}
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word)==0{
		return false
	}
	ary:=this.V
	var tp *t
	for i:=0;i<len(word);i++{
		tp=ary[word[i]-'a']
		if tp==nil{
			return false
		}
		if i==len(word)-1&&!tp.fg{
			return false
		}
		ary=tp.next
	}
	return true
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix)==0{
		return false
	}
	ary:=this.V
	var tp *t
	for i:=0;i<len(prefix);i++{
		tp=ary[prefix[i]-'a']
		if tp==nil{
			return false
		}
		ary=tp.next
	}

	return true
}


