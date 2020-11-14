package pos

import (
	"sort"
)

type MemoSystem struct {
	list [][]event
}
type event struct {
	content string
	finish  bool
}

func Constructor_1() MemoSystem {
	return MemoSystem{
		make([][]event, 1001),
	}
}

func (memoSystem *MemoSystem) checkHas(idx int, content string) bool {
	for j := 0; j < len(memoSystem.list[idx]); j++ {
		if content == memoSystem.list[idx][j].content {
			return true
		}
	}
	return false
}

func (memoSystem *MemoSystem) addEvent(startDate int, content string, num int, period int) int {
	sum := num
	for i := 0; i < num; i++ {
		idx := startDate + i*period
		if memoSystem.checkHas(idx, content) {
			sum--
			continue
		}
		memoSystem.list[idx] = append(memoSystem.list[idx], event{
			content: content,
			finish:  false,
		})
	}
	return sum
}

func (memoSystem *MemoSystem) finishEvent(date int, content string) bool {
	curList := memoSystem.list[date]
	for i := range curList {
		if curList[i].content == content {
			if curList[i].finish {
				return false
			}
			curList[i].finish = true
			return true
		}
	}
	return false
}

func (memoSystem *MemoSystem) removeEvent(date int, content string) bool {
	curList := memoSystem.list[date]
	for i := range curList {
		if curList[i].content == content {
			memoSystem.list[date] = append(memoSystem.list[date][:i], memoSystem.list[date][i+1:]...)
			return true
		}
	}
	return false
}

func (memoSystem *MemoSystem) queryTodo(startDate int, endDate int) []string {
	res := []string{}
	for i := startDate; i <= endDate; i++ {
		curContent := []string{}
		for j := 0; j < len(memoSystem.list[i]); j++ {
			if !memoSystem.list[i][j].finish {
				curContent = append(curContent, memoSystem.list[i][j].content)
			}
		}
		sort.Strings(curContent)
		res = append(res, curContent...)
	}
	return res
}
