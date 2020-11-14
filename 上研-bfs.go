package pos

// bfs
// 1、点的结构
// 2、起点
// 3、终点
// 4、点与点之间的关系
// 5、遍历状态表
// 6、终止条件
// 7、返回结果
var (
	// 迷宫初始化 -1是墙 0不是
	g_input = [100][100]int{}
	// 迷宫大小
	g_size = 100
	// 起点
	g_start = [2]int{0, 0}
	// 终点
	g_end = [2]int{3, 3}
	// 4个方向
	g_dir = [][2]int{{1,0}, {-1,0}, {0,1}, {0, -1}}
	// 队列
	queue = [250000][2]int{}

)
// 判断是否在迷宫内
func isVaild(x int, y int)bool{
	if x<0|| x>g_size|| y<0|| y>g_size{
		return false
	}
	return true
}
// 经典解法
func cals()int{
	// 队列初始化
	begin, end := 0, 0
	// 起点入队
	queue[end][0] = g_start[0]
	queue[end][1] = g_start[1]
	end++
	// 队列判空
	for begin!=end{
		// 坐标出队
		x := queue[begin][0]
		y := queue[begin][1]
		begin++
		// 业务-是否到达终点
		if x==g_end[0] &&y ==g_end[1]{
			return g_input[x][y]
		}
		for i:=0; i<4; i++{
			x1 := x + g_dir[i][0]
			y1 := y + g_dir[i][1]
			// 是否在范围内
			if !isVaild(x1, y1){
				continue
			}
			// 1、是否是墙 2、是否走过
			if g_input[x1][y1]!=0{
				continue
			}
			// 刷新步数
			g_input[x1][y1]++
			// 可走的路径入队
			queue[end][0] = x1
			queue[end][1] = y1
			end++
		}
	}
	return 0
}