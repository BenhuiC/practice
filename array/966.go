package array

import "strings"

func spellchecker(wordlist []string, queries []string) []string {
	wordMatch := make(map[string]bool)
	wordCaseMatch := make(map[string]string)
	wordVowelMatch := make(map[string]string)

	doVoweReplace := func(s string) string {
		bts := []byte(s)
		for i, c := range bts {
			if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
				bts[i] = '*'
			}
		}

		return string(bts)
	}

	for _, w := range wordlist {
		wordMatch[w] = true
		lowCase := strings.ToLower(w)
		if _, ok := wordCaseMatch[lowCase]; !ok {
			wordCaseMatch[lowCase] = w
		}
		voweWord := doVoweReplace(lowCase)
		if _, ok := wordVowelMatch[voweWord]; !ok {
			wordVowelMatch[voweWord] = w
		}
	}

	res := make([]string, 0, len(queries))

	for _, q := range queries {
		if wordMatch[q] {
			res = append(res, q)
			continue
		}
		lowCase := strings.ToLower(q)
		if v, ok := wordCaseMatch[lowCase]; ok {
			res = append(res, v)
			continue
		}
		if v, ok := wordVowelMatch[doVoweReplace(lowCase)]; ok {
			res = append(res, v)
			continue
		}
		res = append(res, "")
	}

	return res
}
