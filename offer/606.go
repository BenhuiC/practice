package offer

import (
	"fmt"
	"strconv"
)

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	l := tree2str(root.Left)
	r := tree2str(root.Right)
	res := strconv.Itoa(root.Val)
	if l == "" && r == "" {
		return res
	} else if r != "" {
		res += fmt.Sprintf("(%s)(%s)", l, r)
	} else {
		res += fmt.Sprintf("(%s)", l)
	}
	return res
}
