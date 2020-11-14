package pos

func backspaceCompare(S string, T string) bool {
	if formateBackspace(S) == formateBackspace(T) {
		return true
	}
	return false
}
func formateBackspace(s string) string {
	ret := ""
	for _, v := range s {
		if len(ret) != 0 && v == '#' {
			ret = ret[:len(ret)-1]
		}
		if v != '#' {
			ret += string(v)
		}
	}
	return ret
}
