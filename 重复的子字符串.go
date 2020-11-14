package pos

import "strings"

func repeatedSubstringPattern(s string) bool {
	res := s + s
	res = res[1 : len(res)-1]
	return strings.Index(res, s) != -1
}
