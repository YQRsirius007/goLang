package pos

import "strings"

func permutation(S string) []string {
	res := make([]string, 0)
	if len(S) == 0 {
		return res
	}
	dfs_permutation(S, &res, "")
	return res
}
func dfs_permutation(s string, res *[]string, path string) {
	if len(path) == len(s) {
		*res = append(*res, path)
	}
	for _, c := range s {
		if !strings.Contains(path, string(c)) {
			dfs_permutation(s, res, path+string(c))
		}
	}
}
