package pos

type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// func minimumLengthEncoding(words []string) int {
// 	sort.Sort((words))
// 	str:=""
// 	for i:= len(words)-1; i>=0; i--{
// 		if !strings.Contains(str, words[i]+"#"){
// 			str +=words[i]+"#"
// 		}
// 	}
// 	return len(str)
// }
