package pos

// 初始化：每个元素所在集合初始化为自身
// 合并：两个元素所属的结合合并为一个集合
// 查找：查找元素所在的集合，即根节点

// Kruskal算法
// 图的连通分量
// 静态连通性
// 动态连通性：判断图的最早连通时间

// 初始化使用数组 维护一个数组存放各个节点的根节点是谁

var pre = [3]int{1, 2, 3}

func union(x, y int) {
	px, py := pre[x], pre[y]
	if px != py {
		px = py
	}
}
func find(x int) int {
	for x != pre[x] {
		// x = pre[x]
		pre[x] = find(pre[x])
	}
	return x
}
