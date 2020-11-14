package pos

import "math/rand"

// 想要0（1） 数组、哈希表
// 采用哈希表去存 元素对应的索引 一步到位
// 重点是设计思路
// 添加元素 在最后添加
// 删除元素 最后删除
type RandomizedSet struct {
	rSclice []int
	rMap    map[int]int
}

func Constructor2() RandomizedSet {
	return RandomizedSet{
		[]int{},
		map[int]int{},
	}
}
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.rMap[val]; ok {
		return false
	}
	this.rSclice = append(this.rSclice, val)
	this.rMap[val] = len(this.rSclice) - 1
	return true
}
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.rMap[val]; !ok {
		return false
	}
	idx := this.rMap[val]
	this.rSclice[idx] = this.rSclice[len(this.rSclice)-1]
	this.rSclice = this.rSclice[:len(this.rSclice)-1]
	this.rMap[this.rSclice[idx]] = idx
	delete(this.rMap, val)
	return true
}
func (this *RandomizedSet) GetRandom() int {
	if len(this.rSclice) == 0 {
		return -1
	}
	idx := rand.Intn(len(this.rSclice))
	return this.rSclice[idx]
}
