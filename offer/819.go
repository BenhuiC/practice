package offer

import "unicode"

func mostCommonWord(paragraph string, banned []string) string {
	banMap := make(map[string]bool)
	for _, w := range banned {
		banMap[w] = true
	}
	freq := make(map[string]int)
	word := make([]byte, 0)
	for i, n := 0, len(paragraph); i <= n; i++ {
		if i < n && unicode.IsLetter(rune(paragraph[i])) {
			word = append(word, byte(unicode.ToLower(rune(paragraph[i]))))
		} else if len(word) > 0 {
			s := string(word)
			if !banMap[s] {
				freq[s]++
			}
			word = word[:0]
		}
	}
	maxFreq := 0
	res := ""
	for k, v := range freq {
		if v > maxFreq {
			res = k
			maxFreq = v
		}
	}
	return res
}
