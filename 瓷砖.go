package pos

func leastBricks(wall [][]int) int {
	if wall == nil || len(wall) == 0 {
		return 0
	}
	row := len(wall)
	interval := make(map[int]int, 0)
	minInterval := 0
	for i := 0; i < row; i++ {
		cur := 0
		for j := 0; j < len(wall[i])-1; j++ {
			cur += wall[i][j]
			_, ok := interval[cur]
			if ok {
				interval[cur]--
			} else {
				interval[cur] = -1
			}
			minInterval = min(interval[cur], minInterval)
		}
	}
	return minInterval + row
}

// func leastBricks(wall [][]int) int{
// 	if wall == nil || len(wall)==0{
// 		return 0
// 	}
// 	row := len(wall)
// 	minInterval, width := 0, 0
// 	for _, v := range wall[0]{
// 		width += v
// 	}
// 	interval := make([]int, width)
// 	for i:=0; i<row; i++{
// 		cur := 0
// 		for j:=0; j<len(wall[i])-1; j++{
// 			cur += wall[i][j]
// 			interval[cur] --
// 			minInterval = min(interval[cur], minInterval)
// 		}
// 	}
// 	minInterval +=  row
// 	return minInterval
// }
// func min(a, b int)int{
// 	if b<a{
// 		return b
// 	}
// 	return a
// }
