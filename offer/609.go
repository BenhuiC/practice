package offer

import "strings"

func findDuplicate(paths []string) [][]string {
	var filesDetail func(s string) (files []string, contents []string)
	filesDetail = func(s string) (files []string, contents []string) {
		ary := strings.Split(s, " ")
		dir := ary[0]
		files, contents = make([]string, 0), make([]string, 0)
		for i := 1; i < len(ary); i++ {
			for j := 0; j < len(ary[i]); j++ {
				if ary[i][j] == '(' {
					files = append(files, dir+"/"+ary[i][:j])
					contents = append(contents, ary[i][j+1:len(ary[i])-1])
				}
			}
		}
		return
	}
	contentMap := make(map[string]string)
	idxMap := make(map[string]int)
	res := make([][]string, 0)
	for _, p := range paths {
		files, contents := filesDetail(p)
		for i, c := range contents {
			if v, ok := contentMap[c]; !ok {
				contentMap[c] = files[i]
			} else {
				idx, ok := idxMap[c]
				if !ok {
					res = append(res, []string{v, files[i]})
					idxMap[c] = len(res) - 1
				} else {
					res[idx] = append(res[idx], files[i])
				}
			}
		}
	}

	return res
}
