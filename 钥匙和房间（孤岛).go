package pos

func canVisitAllRooms(rooms [][]int) bool {
	if len(rooms) == 0 {
		return false
	}
	if len(rooms) == 1 {
		return true
	}
	visited := make([]bool, len(rooms))
	dfs_room(0, rooms, &visited)
	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}
func dfs_room(index int, rooms [][]int, visited *[]bool) {
	(*visited)[index] = true
	for _, v := range rooms[index] {
		dfs_room(v, rooms, visited)
	}
}
