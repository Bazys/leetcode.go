package leetcode

type MyQueue struct {
	stack []int
}

func NewMyQueue() MyQueue {
	return MyQueue{
		stack: make([]int, 0),
	}
}

func (s *MyQueue) Push(x int) {
	s.stack = append(s.stack, x)
}

func (s *MyQueue) Pop() int {
	var x int
	x, s.stack = s.stack[0], s.stack[1:]
	return x
}

func (s *MyQueue) Peek() int {
	return s.stack[0]
}

func (s *MyQueue) Empty() bool {
	if len(s.stack) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
