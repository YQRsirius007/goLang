package pos

import (
	"container/heap"
)

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}
	frequent := map[int]int{}
	for _, v := range nums {
		frequent[v]++
	}

	h := &IHeap{}
	heap.Init(h)
	for key, value := range frequent {
		heap.Push(h, [2]int{key, value})
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[i] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] > h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
