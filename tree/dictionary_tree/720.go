package dictionary_tree

import "sort"

func longestWord(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		s, t := words[i], words[j]
		return len(s) < len(t) || len(s) == len(t) && s > t
	})
	candidates := map[string]struct{}{"": {}}
	res := ""
	for _, w := range words {
		if _, ok := candidates[w[:len(w)-1]]; ok {
			res = w
			candidates[w] = struct{}{}
		}
	}
	return res
}

func longestWord2(words []string) string {
	type dictionaryTree struct {
		end      bool
		children map[byte]*dictionaryTree
	}
	add := func(d *dictionaryTree, word string) int {
		endCount := 0
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
			if c.end {
				endCount++
			}
			cur = c
		}
		return endCount
	}
	d := &dictionaryTree{children: make(map[byte]*dictionaryTree)}
	sort.Slice(words, func(i, j int) bool {
		s, t := words[i], words[j]
		return len(s) < len(t) || len(s) == len(t) && s < t
	})

	res := ""
	for _, w := range words {
		cnt := add(d, w)
		if cnt == len(w) && cnt > len(res) {
			res = w
		}
	}

	return res
}
