package partice

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//示例：
//
//输入：n = 3
//输出：[
//       "((()))",
//       "(()())",
//       "(())()",
//       "()(())",
//       "()()()"
//     ]

var genResult []string

func generateParenthesis(n int) []string {
	genResult = make([]string, 0)
	if n <= 0 {
		return genResult
	}
	gen("", n, 0, 0)
	return genResult
}

// 回溯
func gen(s string, n, l, f int) {
	if l == f && n == l {
		genResult = append(genResult, s)
		return
	}
	if f > l || l > n {
		return
	}
	gen(s+"(", n, l+1, f)
	gen(s+")", n, l, f+1)

	return
}
