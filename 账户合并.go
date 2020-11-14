package pos

func accountsMerge(accounts [][]string) [][]string {
	ret := [][]string{}
	ret = append(ret, accounts[0])
	m := map[string]int{}
	// 根节点
	parent := make([]int, len(accounts))
	for i := range parent {
		parent[i] = i
	}
	for key, account := range accounts {
		for i := 1; i < len(account); i++ {
			if _, ok := m[account[i]]; ok {
				// 已经存在，应该合并
				union1(parent, m[account[i]], key)

				continue
			} else {
				m[account[i]] = key
			}
		}
	}

	return ret
}
func union1(parent []int, x int, y int) {
	if find1(parent, x) != find1(parent, y) {
		parent[y] = parent[x]
	}
}
func find1(parent []int, x int) int {
	for parent[x] != x {
		x = parent[x]
	}
	return x
}
