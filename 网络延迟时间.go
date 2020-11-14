package pos

import (
	"fmt"
	"math"
)

func networkDelayTime(times [][]int, N int, K int) int {
	// 邻接表 存放各个点到各个点的位置
	graph := make([][]int, N+1)
	for i := 1; i < N+1; i++ {
		graph[i] = make([]int, N+1)
		for j := 1; j < N+1; j++ {
			graph[i][j] = -1
		}
	}

	// 遍历times填充邻接表
	for _, time := range times {
		graph[time[0]][time[1]] = time[2]
	}
	fmt.Println(graph)
	// K是源点，存放K到各个点的距离
	distance := make([]int, N+1)
	for i := range distance {
		distance[i] = -1
	}

	// 初始化distance为K到各个节点的距离
	for i := 1; i < N+1; i++ {
		distance[i] = graph[K][i]
	}
	fmt.Println(distance)
	// K到K的节点初始化为0
	distance[K] = 0
	// 判断是否找到K到达该点最短路径
	visited := make([]bool, N+1)
	visited[K] = true

	// 遍历除k本身节点之外的所有N-1个节点
	for i := 1; i < N; i++ {
		minDis := math.MaxInt32
		minIndex := 1
		// 离K最近的节点
		for j := 1; j < N+1; j++ {
			if !visited[j] && distance[j] != -1 && distance[j] < minDis {
				minIndex = j
				minDis = distance[j]
			}
		}
		fmt.Println("minIndex", minIndex)
		visited[minIndex] = true
		// 更新k与其他节点的距离
		for j := 1; j < N+1; j++ {
			if graph[minIndex][j] != -1 {
				if distance[j] == -1 {
					distance[j] = distance[minIndex] + graph[minIndex][j]
				} else {
					distance[j] = min(distance[j], distance[minIndex]+graph[minIndex][j])
				}
			}
		}
	}
	maxDistance := 0
	for i := 1; i < N+1; i++ {
		if distance[i] == -1 {
			return -1
		}
		maxDistance = max(distance[i], maxDistance)
	}
	return maxDistance
}

// func dfs_network(times [][]int, marked *map[int]int, start int, maxPath *int){
// 	for i, time := range times{
// 		if time[0]>start{
// 			break
// 		}
// 		if time[0]==start{
// 			(*marked)[time[1]] = 1
// 			dfs_network(times, marked, times[1], maxPath)
// 		}
// 	}
// }
