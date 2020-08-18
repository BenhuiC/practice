package partice

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func longestPalindrome(s string) string {
	result := ""
	// 长度小于2即可返回原字符串
	if len(s) <= 1 {
		return s
	}
	for i := 0; i < len(s); i++ {
		i = longestStr(s, i, &result)
	}

	return result
}

func longestStr(str string, start int, result *string) int {
	var end, re int
	for end = start; end < len(str)-1 && str[end+1] == str[start]; end++ {
	}
	re = end
	for start > 0 && end < len(str)-1 && str[end+1] == str[start-1] {
		start--
		end++
	}
	if end-start+1 > len(*result) {
		*result = str[start : end+1]
	}
	return re
}
