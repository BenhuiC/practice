package partice

var morse = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	m := make(map[string]struct{})
	for _, v := range words {
		t := ""
		for _, c := range v {
			t += morse[c-'a']
		}
		m[t] = struct{}{}
	}

	return len(m)
}
