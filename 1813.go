package partice

import "strings"

/*
相似的情况有三种
1.短字符串是长字符串的前缀；
2.短字符串是长字符串的后缀；
3.短字符串前x个单词与长字符串的前x个单词相同，后y个单词和长字符串的后y个单词相同；
思路
1.计算两个字符串的首部有left个单词相同
2.计算两个字符串的尾部有right个单词相同
3.对于情况1:则left等于短字符串的单词数量，即left+right>=len(short)；情况2与1类似；情况3:首部相同的字符串数量加上尾部字符串相同的数量和必大于短字符串的单词数，即left+right>=len(short)
4.所以求前left个单词相同，后right个单词相同，left+right>=len(short)的情况即为相似
*/

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	long := strings.Split(sentence1, " ")
	short := strings.Split(sentence2, " ")
	if len(short) > len(long) {
		long, short = short, long
	}
	left, right := 0, 0 // 分别从前后遍历，相同的单词数
	// 从前遍历
	for i := 0; i < len(long) && i < len(short); i++ {
		if long[i] != short[i] {
			break
		}
		left++
	}
	// 从尾部遍历
	j, x := len(long)-1, len(short)-1
	for j >= 0 && x >= 0 {
		if long[j] != short[x] {
			break
		}
		right++
		j--
		x--
	}
	if left+right >= len(short) {
		return true
	}
	return false
}
