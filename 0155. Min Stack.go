package leetcode

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int             { return len(h) }
func (h IntHeap) Less(i, j int) bool   { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)        { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{})  { *h = append(*h, x.(int)) }
func (h IntHeap) At(i int) interface{} { return h[i] }
func (h *IntHeap) Pop() (x interface{}) {
	old := *h
	x, *h = old[len(old)-1], old[0:len(old)-1]
	return
}

type MinStack struct {
	deleted map[int]struct{}
	stack   []int
	queue   *IntHeap
}

func NewMinStack() MinStack {
	queue := &IntHeap{}
	heap.Init(queue)
	return MinStack{
		deleted: map[int]struct{}{},
		stack:   []int{},
		queue:   queue,
	}
}

func (s *MinStack) Push(val int) {
	s.stack = append(s.stack, val)
	heap.Push(s.queue, val)

}

func (s *MinStack) Pop() {
	var x int
	x, s.stack = s.stack[len(s.stack)-1], s.stack[:len(s.stack)-1]
	s.deleted[x] = struct{}{}
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	min := s.queue.At(0).(int)
	for {
		if _, ok := s.deleted[min]; ok {
			heap.Pop(s.queue)
			delete(s.deleted, min)
			min = s.queue.At(0).(int)
			continue
		}
		return s.queue.At(0).(int)
	}
}
