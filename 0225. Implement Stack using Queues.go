package leetcode

type MyStack struct {
	stack []int
}

func NewMyStack() MyStack {
	return MyStack{
		stack: make([]int, 0),
	}
}

func (s *MyStack) Push(x int) {
	s.stack = append(s.stack, x)
}

func (s *MyStack) Pop() int {
	var x int
	n := len(s.stack)
	x, s.stack = s.stack[n-1], s.stack[:n-1]
	return x
}

func (s *MyStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MyStack) Empty() bool {
	if len(s.stack) == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
