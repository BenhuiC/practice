package offer

import (
	"fmt"
	"testing"
)

func TestWordDictionary_AddWord(t *testing.T) {
	w := Constructor()
	w.AddWord("abcd")
	w.AddWord("adce")
	w.AddWord("a")
	fmt.Println(w.Search("ab.."))
}
