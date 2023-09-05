package types

import "strings"

func replaceWords(dictionary []string, sentence string) string {
	if len(dictionary) == 0 {
		return sentence
	}

	t := &trieTree{child: [26]*trieTree{}}
	for _, w := range dictionary {
		t.insert(w)
	}

	arys := strings.Split(sentence, " ")
	for i, str := range arys {
		toReplace := []byte{}

		tmp := t
		for _, s := range str {
			if tmp.child[s-'a'] == nil {
				break
			}
			toReplace = append(toReplace, byte(s))
			tmp = tmp.child[s-'a']
			if tmp.end {
				if len(toReplace) != 0 {
					arys[i] = string(toReplace)
				}
				break
			}
		}
	}

	return strings.Join(arys, " ")
}

type trieTree struct {
	data  byte
	end   bool
	child [26]*trieTree
}

func (t *trieTree) insert(word string) {
	p := t
	for _, v := range word {
		if p.child[v-'a'] == nil {
			p.child[v-'a'] = &trieTree{
				data:  byte(v),
				child: [26]*trieTree{},
			}
		}
		p = p.child[v-'a']
	}
	p.end = true
}
