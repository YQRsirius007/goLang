package pos
// 细节是魔鬼
//  1、 判断条件是否是动态变量 需不需要用临时变量缓存
// 2、判断条件是否有等于
func ladderLength(beginWord string, endWord string, wordList[]string)int{
	wordMap := map[string]bool{}
	for _, word := range wordList{
		wordMap[word] = true
	}
	var queue []string
	queue = append(queue, beginWord)
	level := 1
	// vistied[beginWord] = 1
	for len(queue)!=0{
		size := len(queue)
		for i:=0; i<size; i++{
			word := queue[0]
			queue = queue[1:]
			if word == endWord{
				return level
			}
			for j:=0; j<len(word); j++{
				for c:='a'; c<='z';c++{
					newWord := word[:j]+string(c)+word[j+1:]
					if wordMap[newWord]{
						queue = append(queue, newWord)
						wordMap[newWord] = false
					}
				}
			}
		}
		level++
	}
	return 0
}
// func ladderLength2(beginWord string, endWord string, wordList[]string)int {
// 	vistied := map[string]bool{}
// 	var queue []string
// 	for _, word := range wordList{
// 		vistied[word] = false
// 	}
// 	if _ ,ok := vistied[endWord]; !ok{
// 		return 0
// 	}
// 	queue = append(queue, beginWord)
// 	count := 0
// 	for len(queue)!=0{
// 		size :=len(queue)
// 		count++
// 		for i:=0; i<size; i++{
// 			start := queue[0]
// 			queue = queue[1:]
// 			for
// 		}
// 	}
// }

